package level2

import (
	"encoding/json"

	"github.com/JetBlink/order_book/helper"

	"github.com/JetBlink/order_book/order_book"
	"github.com/JetBlink/order_book/skiplist"
	"github.com/shopspring/decimal"
)

type OrderBook struct {
	*order_book.OrderBook
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		order_book.NewOrderBook(newAskOrders(), newBidOrders()),
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

func (fb *OrderBook) addToOrderBookSide(book *skiplist.SkipList, order *Order) {
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

func (fb *OrderBook) AddAsk(order *Order) {
	fb.addToOrderBookSide(fb.Asks, order)
}

func (fb *OrderBook) AddBid(order *Order) {
	fb.addToOrderBookSide(fb.Bids, order)
}

func (fb *OrderBook) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"sequence":         fb.Sequence,
		order_book.AskSide: fb.GetPartOrderBookBySide(order_book.AskSide, 0),
		order_book.BidSide: fb.GetPartOrderBookBySide(order_book.BidSide, 0),
	})
}

func (fb *OrderBook) GetPartOrderBookBySide(side string, number int) [][2]string {
	var it skiplist.Iterator
	if side == order_book.AskSide {
		it = fb.Asks.Iterator()
		number = helper.Min(number, fb.Asks.Len())
		if number == 0 {
			number = fb.Asks.Len()
		}
	} else {
		it = fb.Bids.Iterator()
		number = helper.Min(number, fb.Bids.Len())

		if number == 0 {
			number = fb.Bids.Len()
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
