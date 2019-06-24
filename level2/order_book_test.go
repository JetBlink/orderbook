package level2

import (
	"log"
	"testing"

	"github.com/JetBlink/order_book/helper"
)

func TestOrderBook_AddAsk(t *testing.T) {
	orderBook := NewOrderBook()
	order, err := NewOrder("10", "20")
	if err != nil {
		t.Error(err)
	}

	orderBook.AddBid(order)

	log.Println(helper.ToJsonString(orderBook))
}
