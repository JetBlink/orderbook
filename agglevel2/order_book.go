package agglevel2

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/JetBlink/orderbook/base"
	"github.com/JetBlink/orderbook/skiplist"
	"github.com/shopspring/decimal"
)

// aggregated_l2 book
type OrderBook struct {
	Asks       *skiplist.SkipList //ask 是 要价，喊价 卖家 卖单 Sort price from low to high
	Bids       *skiplist.SkipList //bid 是 投标，买家 买单 Sort price from high to low
	tickerSize decimal.Decimal
	//聚合前订单
	asksOrderPool map[string]*Order
	bidsOrderPool map[string]*Order
}

// tickerSize like 0.05
func NewOrderBook(tickerSize string) *OrderBook {
	return &OrderBook{
		newAskOrders(),
		newBidOrders(),
		decimal.RequireFromString(tickerSize),
		make(map[string]*Order),
		make(map[string]*Order),
	}
}

func isEqual(l, r interface{}) bool {
	switch val := l.(type) {
	case decimal.Decimal:
		cVal := r.(decimal.Decimal)
		if !val.Equals(cVal) {
			return false
		}
	default:
		if val != r {
			return false
		}
	}
	return true
}

func newAskOrders() *skiplist.SkipList {
	return skiplist.NewCustomMap(func(l, r interface{}) bool {
		return l.(decimal.Decimal).LessThan(r.(decimal.Decimal))
	}, isEqual)
}

func newBidOrders() *skiplist.SkipList {
	return skiplist.NewCustomMap(func(l, r interface{}) bool {
		return l.(decimal.Decimal).GreaterThan(r.(decimal.Decimal))
	}, isEqual)
}

//获取 ticker size
func (ob *OrderBook) TickerSize() string {
	return ob.tickerSize.String()
}

func (ob *OrderBook) getOrderBookBySide(side string) (*skiplist.SkipList, error) {
	if err := base.CheckSide(side); err != nil {
		return nil, err
	}

	if side == base.AskSide {
		return ob.Asks, nil
	}

	return ob.Bids, nil
}

func (ob *OrderBook) getOrderPoolBySide(side string) (map[string]*Order, error) {
	if err := base.CheckSide(side); err != nil {
		return nil, err
	}

	if side == base.AskSide {
		return ob.asksOrderPool, nil
	}

	return ob.bidsOrderPool, nil
}

func (ob *OrderBook) getAggregatedPrice(side string, price decimal.Decimal) (fixPrice decimal.Decimal, err error) {
	mod := price.Mod(ob.tickerSize)

	switch side {
	case base.AskSide:
		if mod.IsZero() {
			fixPrice = price
		} else {
			fixPrice = price.Sub(mod).Add(ob.tickerSize)
		}
	case base.BidSide:
		fixPrice = price.Sub(mod)
	default:
		err = errors.New(fmt.Sprintf("error side, side should be %s or %s", base.AskSide, base.BidSide))
		return
	}

	return
}

//add or set To Order Book Side
func (ob *OrderBook) SetOrder(side string, order *Order) error {
	orderBook, err := ob.getOrderBookBySide(side)
	if err != nil {
		return err
	}
	orderPool, err := ob.getOrderPoolBySide(side)
	if err != nil {
		return err
	}
	aggregatedPrice, err := ob.getAggregatedPrice(side, order.Price)
	if err != nil {
		return err
	}
	node, aggregatedPriceExist := orderBook.GetNode(aggregatedPrice)
	oldOrder, oldOrderExist := orderPool[order.Price.String()]

	//删除
	if order.Size.IsZero() {
		if !oldOrderExist {
			return nil
			//return errors.New("不应该出现")
		}
		if aggregatedPriceExist {
			value := node.Get().(*aggregatedOrder)
			value.Size = value.Size.Sub(oldOrder.Size)
			if value.Size.LessThan(decimal.Zero) {
				panic("size 变负数了 1")
			}
			if value.Size.IsZero() {
				orderBook.Delete(aggregatedPrice)
			}
		}
		delete(orderPool, order.Price.String())
		return nil
	}

	//新增
	if !aggregatedPriceExist {
		if oldOrderExist {
			return errors.New("不应该出现")
		} else {
			//fmt.Println("xxxx", aggregatedPrice, order.Size)

			orderBook.Set(aggregatedPrice, &aggregatedOrder{
				AggregatedPrice: aggregatedPrice,
				Size:            order.Size,
			})
			orderPool[order.Price.String()] = order
		}
	} else { //聚合价格存在
		value := node.Get().(*aggregatedOrder)
		if !oldOrderExist {
			value.Size = value.Size.Add(order.Size)
			orderPool[order.Price.String()] = order
		} else {
			value.Size = value.Size.Sub(oldOrder.Size.Sub(order.Size))
			if value.Size.LessThan(decimal.Zero) {
				panic("size 变负数了 2")
			}
			if value.Size.IsZero() {
				return errors.New("不应该存在")
			}
			oldOrder.Size = order.Size
		}
	}

	return nil
}

func (ob *OrderBook) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		base.AskSide: ob.GetPartOrderBookBySide(base.AskSide, 0),
		base.BidSide: ob.GetPartOrderBookBySide(base.BidSide, 0),
	})
}

func (ob *OrderBook) GetPartOrderBookBySide(side string, number int) [][2]string {
	if err := base.CheckSide(side); err != nil {
		return nil
	}

	var it skiplist.Iterator
	if side == base.AskSide {
		it = ob.Asks.Iterator()
		if number == 0 {
			number = ob.Asks.Len()
		} else {
			number = base.Min(number, ob.Asks.Len())
		}
	} else {
		it = ob.Bids.Iterator()
		if number == 0 {
			number = ob.Bids.Len()
		} else {
			number = base.Min(number, ob.Bids.Len())
		}
	}

	arr := make([][2]string, number)
	it.Next()

	for i := 0; i < number; i++ {
		order := it.Value().(*aggregatedOrder)
		arr[i] = [2]string{order.AggregatedPrice.String(), order.Size.String()}
		if !it.Next() {
			break
		}
	}

	return arr
}
