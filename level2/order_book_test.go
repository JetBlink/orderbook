package level2

import (
	"log"
	"testing"

	"github.com/JetBlink/order_book/helper"
)

func TestOrderBook_SetOrder(t *testing.T) {
	orderBook := NewOrderBook()
	order, err := NewOrder("10", "20")
	if err != nil {
		t.Error(err)
	}

	orderBook.SetOrder(AskSide, order)

	log.Println("TestOrderBook_SetOrder", helper.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrderWithSize0(t *testing.T) {
	orderBook := NewOrderBook()
	order, _ := NewOrder("10", "1.0")
	orderBook.SetOrder(AskSide, order)

	order0, _ := NewOrder("10", "0")
	orderBook.SetOrder(AskSide, order0)

	log.Println("TestOrderBook_SetOrderWithSize0", helper.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 0 {
		t.Error("ask len is not 0")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrderWithBidSize0(t *testing.T) {
	orderBook := NewOrderBook()
	order, _ := NewOrder("10", "1.0")
	orderBook.SetOrder(AskSide, order)

	order0, _ := NewOrder("10", "0")
	orderBook.SetOrder(BidSide, order0)

	log.Println("TestOrderBook_SetOrderWithSize0", helper.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrderWithSamePrice(t *testing.T) {
	orderBook := NewOrderBook()
	order, _ := NewOrder("10", "1.0")
	orderBook.SetOrder(AskSide, order)

	order0, _ := NewOrder("10", "2.2")
	orderBook.SetOrder(AskSide, order0)

	log.Println("TestOrderBook_SetOrderWithSamePrice", helper.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrder2(t *testing.T) {
	orderBook := NewOrderBook()
	order, _ := NewOrder("10", "1.0")
	orderBook.SetOrder(AskSide, order)

	order0, _ := NewOrder("10", "2.2")
	orderBook.SetOrder(BidSide, order0)

	log.Println("TestOrderBook_SetOrderWithSamePrice", helper.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 1 {
		t.Error("bid len is not 1")
	}
}

func TestOrderBook_SetOrder3(t *testing.T) {
	orderBook := NewOrderBook()
	order, _ := NewOrder("0.1", "1.0")
	orderBook.SetOrder(AskSide, order)

	order1, _ := NewOrder("0.3", "1.1")
	orderBook.SetOrder(AskSide, order1)

	order2, _ := NewOrder("0.2", "2.2")
	orderBook.SetOrder(AskSide, order2)

	order3, _ := NewOrder("0.2", "2.3")
	orderBook.SetOrder(AskSide, order3)

	log.Println("TestOrderBook_SetOrderWithSamePrice", helper.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 3 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}
