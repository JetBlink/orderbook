package level2

import (
	"log"
	"testing"

	"github.com/JetBlink/orderbook/base"
)

func TestOrderBook_SetOrder(t *testing.T) {
	orderBook := NewOrderBook()
	order, err := NewOrder("10", "20", nil)
	if err != nil {
		t.Error(err)
	}

	orderBook.SetOrder(base.AskSide, order)

	log.Println("TestOrderBook_SetOrder", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrderWithSize0(t *testing.T) {
	orderBook := NewOrderBook()
	order, _ := NewOrder("10", "1.0", nil)
	orderBook.SetOrder(base.AskSide, order)

	order0, _ := NewOrder("10", "0", nil)
	orderBook.SetOrder(base.AskSide, order0)

	log.Println("TestOrderBook_SetOrderWithSize0", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 0 {
		t.Error("ask len is not 0")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrderWithBidSize0(t *testing.T) {
	orderBook := NewOrderBook()
	order, _ := NewOrder("10", "1.0", nil)
	orderBook.SetOrder(base.AskSide, order)

	order0, _ := NewOrder("10", "0", nil)
	orderBook.SetOrder(base.BidSide, order0)

	log.Println("TestOrderBook_SetOrderWithSize0", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrderWithSamePrice(t *testing.T) {
	orderBook := NewOrderBook()
	order, _ := NewOrder("10", "1.0", nil)
	orderBook.SetOrder(base.AskSide, order)

	order0, _ := NewOrder("10", "2.2", nil)
	orderBook.SetOrder(base.AskSide, order0)

	log.Println("TestOrderBook_SetOrderWithSamePrice", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrder2(t *testing.T) {
	orderBook := NewOrderBook()
	order, _ := NewOrder("10", "1.0", nil)
	orderBook.SetOrder(base.AskSide, order)

	order0, _ := NewOrder("10", "2.2", nil)
	orderBook.SetOrder(base.BidSide, order0)

	log.Println("TestOrderBook_SetOrderWithSamePrice", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 1 {
		t.Error("bid len is not 1")
	}
}

func TestOrderBook_SetOrder3(t *testing.T) {
	orderBook := NewOrderBook()
	order, _ := NewOrder("0.1", "1.0", nil)
	orderBook.SetOrder(base.AskSide, order)

	order1, _ := NewOrder("0.3", "1.1", nil)
	orderBook.SetOrder(base.AskSide, order1)

	order2, _ := NewOrder("0.2", "2.2", nil)
	orderBook.SetOrder(base.AskSide, order2)

	order3, _ := NewOrder("0.2", "2.3", nil)
	orderBook.SetOrder(base.AskSide, order3)

	log.Println("TestOrderBook_SetOrderWithSamePrice", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 3 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}
