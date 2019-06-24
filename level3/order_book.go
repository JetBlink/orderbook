package level3

import (
	"encoding/json"
	"fmt"

	"github.com/JetBlink/order_book/helper"
	"github.com/JetBlink/order_book/skiplist"
	"github.com/shopspring/decimal"
)

const (
	AskSide = "asks"
	BidSide = "bids"
)

type OrderBook struct {
	Sequence  uint64             //Sequence || UpdateID
	Asks      *skiplist.SkipList //ask 是 要价，喊价 卖家 卖单 Sort price from low to high
	Bids      *skiplist.SkipList //bid 是 投标，买家 买单 Sort price from high to low
	OrderPool map[string]*Order
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		Sequence:  0,
		Asks:      newAskOrders(),
		Bids:      newBidOrders(),
		OrderPool: make(map[string]*Order),
	}
}

func isEqual(l, r interface{}) bool {
	switch val := l.(type) {
	case decimal.Decimal:
		cVal := r.(decimal.Decimal)
		if !val.Equals(cVal) {
			return false
		}

	case *Order:
		cVal := r.(*Order)
		if cVal.OrderId != val.OrderId {
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
		if l.(*Order).Price.Equal(r.(*Order).Price) {
			return l.(*Order).Time < r.(*Order).Time
		}

		return l.(*Order).Price.LessThan(r.(*Order).Price)
	}, isEqual)
}

func newBidOrders() *skiplist.SkipList {
	return skiplist.NewCustomMap(func(l, r interface{}) bool {
		if l.(*Order).Price.Equal(r.(*Order).Price) {
			return l.(*Order).Time < r.(*Order).Time
		}

		return l.(*Order).Price.GreaterThan(r.(*Order).Price)
	}, isEqual)
}

func (ob *OrderBook) getOrderBookBySide(side string) *skiplist.SkipList {
	if side == AskSide {
		return ob.Asks
	}

	return ob.Bids
}

//open 事件
func (ob *OrderBook) AddOrder(side string, order *Order) {
	ob.getOrderBookBySide(side).Set(order, order)
	ob.OrderPool[order.OrderId] = order
}

//done 事件
func (ob *OrderBook) RemoveByOrderId(side string, orderId string) {
	order, ok := ob.OrderPool[orderId]
	if !ok {
		return
	}

	ob.removeOrder(side, order)
}

func (ob *OrderBook) removeOrder(side string, order *Order) {
	if _, ok := ob.getOrderBookBySide(side).Delete(order); ok {
		delete(ob.OrderPool, order.OrderId)
	}
}

func (ob *OrderBook) GetOrder(orderId string) *Order {
	order, ok := ob.OrderPool[orderId]
	if !ok {
		return nil
	}

	return order
}

//更新, match 事件
func (ob *OrderBook) MatchOrder(side string, orderId string, size string) error {
	order, ok := ob.OrderPool[orderId]
	if !ok {
		return nil
	}

	sizeValue, err := decimal.NewFromString(size)
	if err != nil {
		return err
	}

	newSize := order.Size.Sub(sizeValue)
	if newSize.LessThan(decimal.Zero) {
		return fmt.Errorf("oldSize: %s, size: %s, sub result less than zero", order.Size.String(), size)
	}

	order.Size = newSize
	if order.Size.Equal(decimal.Zero) {
		ob.removeOrder(side, order)
	}

	return nil
}

//替换, change 事件
func (ob *OrderBook) ChangeOrder(side string, orderId string, size string) error {
	order, ok := ob.OrderPool[orderId]
	if !ok {
		return nil
	}

	sizeValue, err := decimal.NewFromString(size)
	if err != nil {
		return err
	}

	order.Size = sizeValue
	if order.Size.Equal(decimal.Zero) {
		ob.removeOrder(side, order)
	}
	return nil
}

func (ob *OrderBook) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"sequence": ob.Sequence,
		AskSide:    ob.GetPartOrderBookBySide(AskSide, 0),
		BidSide:    ob.GetPartOrderBookBySide(BidSide, 0),
	})
}

func (ob *OrderBook) GetPartOrderBookBySide(side string, number int) [][3]string {
	var it skiplist.Iterator
	if side == AskSide {
		it = ob.Asks.Iterator()
		number = helper.Min(number, ob.Asks.Len())
		if number == 0 {
			number = ob.Asks.Len()
		}
	} else {
		it = ob.Bids.Iterator()
		number = helper.Min(number, ob.Bids.Len())

		if number == 0 {
			number = ob.Bids.Len()
		}
	}

	arr := make([][3]string, number)
	it.Next()

	for i := 0; i < number; i++ {
		order := it.Value().(*Order)
		arr[i] = [3]string{order.OrderId, order.Price.String(), order.Size.String()}
		if !it.Next() {
			break
		}
	}

	return arr
}
