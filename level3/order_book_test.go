package level3

import (
	"log"
	"testing"

	"github.com/JetBlink/order_book/helper"
)

func TestOrderBook_AddAsk(t *testing.T) {
	orderBook := NewOrderBook()
	order, err := NewOrder("abc", "10", "20", uint64(100))
	if err != nil {
		t.Error(err)
	}

	orderBook.AddOrder(AskSide, order)
	orderBook.AddOrder(AskSide, order)

	log.Println(helper.ToJsonString(orderBook))
}
