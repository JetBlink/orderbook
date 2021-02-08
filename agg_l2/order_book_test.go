package agg_l2

import (
	"log"
	"testing"

	"github.com/JetBlink/orderbook/base"
)

func TestOrderBook_SetOrder(t *testing.T) {
	tickerSize := "0.05"
	orderBook := NewOrderBook(tickerSize)
	order, _ := NewOrder("9.20", "20", nil)
	if err := orderBook.SetOrder(base.AskSide, order); err != nil {
		panic(err)
	}
	order1, _ := NewOrder("9.19", "20", nil)
	if err := orderBook.SetOrder(base.AskSide, order1); err != nil {
		panic(err)
	}

	log.Println("TestOrderBook_SetOrder", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrderWithSize0(t *testing.T) {
	tickerSize := "0.01"
	orderBook := NewOrderBook(tickerSize)
	order, _ := NewOrder("10.10", "20", nil)
	if err := orderBook.SetOrder(base.AskSide, order); err != nil {
		panic(err)
	}

	order0, _ := NewOrder("10.10", "0", nil)
	if err := orderBook.SetOrder(base.AskSide, order0); err != nil {
		panic(err)
	}

	log.Println("TestOrderBook_SetOrderWithSize0", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 0 {
		t.Error("ask len is not 0")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrderWithBidSize0(t *testing.T) {
	tickerSize := "0.01"
	orderBook := NewOrderBook(tickerSize)
	order, _ := NewOrder("10.01", "1.0", nil)
	if err := orderBook.SetOrder(base.AskSide, order); err != nil {
		panic(err)
	}

	order0, _ := NewOrder("10.00", "0", nil)
	if err := orderBook.SetOrder(base.BidSide, order0); err != nil {
		panic(err)
	}

	log.Println("TestOrderBook_SetOrderWithSize0", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrderWithSamePrice(t *testing.T) {
	tickerSize := "0.05"
	orderBook := NewOrderBook(tickerSize)
	order, _ := NewOrder("10.07", "1.0", nil)
	if err := orderBook.SetOrder(base.AskSide, order); err != nil {
		panic(err)
	}

	order0, _ := NewOrder("10.06", "2.2", nil)
	if err := orderBook.SetOrder(base.AskSide, order0); err != nil {
		panic(err)
	}

	log.Println("TestOrderBook_SetOrderWithSamePrice", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrder2(t *testing.T) {
	tickerSize := "0.01"
	orderBook := NewOrderBook(tickerSize)
	order, _ := NewOrder("10.11", "1.0", nil)
	if err := orderBook.SetOrder(base.AskSide, order); err != nil {
		panic(err)
	}

	order0, _ := NewOrder("10.01", "2.2", nil)
	if err := orderBook.SetOrder(base.BidSide, order0); err != nil {
		panic(err)
	}

	log.Println("TestOrderBook_SetOrder2", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 1 {
		t.Error("bid len is not 1")
	}
}

func TestOrderBook_SetOrder3(t *testing.T) {
	tickerSize := "0.5"
	orderBook := NewOrderBook(tickerSize)
	order, _ := NewOrder("10.1", "1.0", nil)
	if err := orderBook.SetOrder(base.AskSide, order); err != nil {
		panic(err)
	}

	order1, _ := NewOrder("10.0", "1.1", nil)
	if err := orderBook.SetOrder(base.AskSide, order1); err != nil {
		panic(err)
	}

	order11, _ := NewOrder("10.0", "0.3", nil)
	if err := orderBook.SetOrder(base.AskSide, order11); err != nil {
		panic(err)
	}

	order111, _ := NewOrder("10.0", "0", nil)
	if err := orderBook.SetOrder(base.AskSide, order111); err != nil {
		panic(err)
	}

	order2, _ := NewOrder("10.3", "2.2", nil)
	if err := orderBook.SetOrder(base.AskSide, order2); err != nil {
		panic(err)
	}

	order3, _ := NewOrder("9.6", "2.3", nil)
	if err := orderBook.SetOrder(base.AskSide, order3); err != nil {
		panic(err)
	}

	order4, _ := NewOrder("9.7", "0", nil)
	if err := orderBook.SetOrder(base.AskSide, order4); err != nil {
		panic(err)
	}

	order5, _ := NewOrder("9.6", "0", nil)
	if err := orderBook.SetOrder(base.AskSide, order5); err != nil {
		panic(err)
	}

	log.Println("TestOrderBook_SetOrder3", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("bid len is not 0")
	}
}

func TestOrderBook_SetOrder4(t *testing.T) {
	tickerSize := "0.5"
	orderBook := NewOrderBook(tickerSize)
	order, _ := NewOrder("10.1", "1.0", nil)
	if err := orderBook.SetOrder(base.BidSide, order); err != nil {
		panic(err)
	}

	order1, _ := NewOrder("10.0", "1.1", nil)
	if err := orderBook.SetOrder(base.BidSide, order1); err != nil {
		panic(err)
	}

	order11, _ := NewOrder("10.0", "0.3", nil)
	if err := orderBook.SetOrder(base.BidSide, order11); err != nil {
		panic(err)
	}

	order111, _ := NewOrder("10.0", "0", nil)
	if err := orderBook.SetOrder(base.BidSide, order111); err != nil {
		panic(err)
	}

	order2, _ := NewOrder("10.3", "2.2", nil)
	if err := orderBook.SetOrder(base.BidSide, order2); err != nil {
		panic(err)
	}

	order3, _ := NewOrder("9.6", "2.3", nil)
	if err := orderBook.SetOrder(base.BidSide, order3); err != nil {
		panic(err)
	}

	order4, _ := NewOrder("9.7", "0", nil)
	if err := orderBook.SetOrder(base.BidSide, order4); err != nil {
		panic(err)
	}

	order5, _ := NewOrder("9.6", "0", nil)
	if err := orderBook.SetOrder(base.BidSide, order5); err != nil {
		panic(err)
	}

	order6, _ := NewOrder("10.5", "1", nil)
	if err := orderBook.SetOrder(base.AskSide, order6); err != nil {
		panic(err)
	}

	order7, _ := NewOrder("10.4", "1", nil)
	if err := orderBook.SetOrder(base.AskSide, order7); err != nil {
		panic(err)
	}

	log.Println("TestOrderBook_SetOrder4", base.ToJsonString(orderBook))
	if orderBook.Asks.Len() != 1 {
		t.Error("ask len is not 1")
	}

	if orderBook.Bids.Len() != 1 {
		t.Error("bid len is not 1")
	}
}
