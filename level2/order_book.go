package level2

import (
	"encoding/json"

	"github.com/JetBlink/orderbook/helper"
	"github.com/JetBlink/orderbook/skiplist"
	"github.com/shopspring/decimal"
)

const (
	AskSide = "asks"
	BidSide = "bids"
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

func (ob *OrderBook) getOrderBookBySide(side string) *skiplist.SkipList {
	if side == AskSide {
		return ob.Asks
	}

	return ob.Bids
}

func (ob *OrderBook) SetOrder(side string, order *Order) {
	ob.addToOrderBookSide(ob.getOrderBookBySide(side), order)
}

func (ob *OrderBook) addToOrderBookSide(book *skiplist.SkipList, order *Order) {
	//if !order.Size.Equal(decimal.Zero) { // New price level
	//	book.Set(order.Price, order)
	//} else if _, ok := book.Get(order.Price); ok { // Existing price level and Quantity Equal 0
	//	book.Delete(order.Price)
	//}

	if order.Size.Equal(decimal.Zero) {
		book.Delete(order.Price)
		return
	}

	book.Set(order.Price, order)
}

func (ob *OrderBook) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"sequence": ob.Sequence,
		AskSide:    ob.GetPartOrderBookBySide(AskSide, 0),
		BidSide:    ob.GetPartOrderBookBySide(BidSide, 0),
	})
}

func (ob *OrderBook) GetPartOrderBookBySide(side string, number int) [][2]string {
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
