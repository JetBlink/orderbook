package order_book

import "github.com/JetBlink/order_book/skiplist"

const (
	AskSide = "asks"
	BidSide = "bids"
)

type OrderBook struct {
	Sequence  uint64             //Sequence || UpdateID
	Asks      *skiplist.SkipList //ask 是 要价，喊价 卖家 卖单 Sort price from low to high
	Bids      *skiplist.SkipList //bid 是 投标，买家 买单 Sort price from high to low
	OrderPool map[string]interface{}
}

func NewOrderBook(Asks, Bids *skiplist.SkipList) *OrderBook {
	return &OrderBook{
		Sequence:  0,
		Asks:      Asks,
		Bids:      Bids,
		OrderPool: make(map[string]interface{}),
	}
}
