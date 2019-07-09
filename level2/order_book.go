package level2

import (
	"encoding/json"

	"github.com/JetBlink/orderbook/base"
	"github.com/JetBlink/orderbook/skiplist"
	"github.com/shopspring/decimal"
)

type OrderBook struct {
	Sequence uint64             //Sequence || UpdateID
	Asks     *skiplist.SkipList //ask 是 要价，喊价 卖家 卖单 Sort price from low to high
	Bids     *skiplist.SkipList //bid 是 投标，买家 买单 Sort price from high to low
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		0,
		newAskOrders(),
		newBidOrders(),
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

func (ob *OrderBook) getOrderBookBySide(side string) (*skiplist.SkipList, error) {
	if err := base.CheckSide(side); err != nil {
		return nil, err
	}

	if side == base.AskSide {
		return ob.Asks, nil
	}

	return ob.Bids, nil
}

//addToOrderBookSide
func (ob *OrderBook) SetOrder(side string, order *Order) error {
	orderBook, err := ob.getOrderBookBySide(side)
	if err != nil {
		return err
	}

	if order.Size.Equal(decimal.Zero) {
		orderBook.Delete(order.Price)
		return nil
	}

	orderBook.Set(order.Price, order)
	return nil
}

func (ob *OrderBook) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"sequence":   ob.Sequence,
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
		number = base.Min(number, ob.Asks.Len())
		if number == 0 {
			number = ob.Asks.Len()
		}
	} else {
		it = ob.Bids.Iterator()
		number = base.Min(number, ob.Bids.Len())

		if number == 0 {
			number = ob.Bids.Len()
		}
	}

	arr := make([][2]string, number)
	it.Next()

	for i := 0; i < number; i++ {
		order := it.Value().(*Order)
		arr[i] = [2]string{order.Price.String(), order.Size.String()}
		if !it.Next() {
			break
		}
	}

	return arr
}
