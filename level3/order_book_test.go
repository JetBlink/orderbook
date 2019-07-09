package level3

import (
	"log"
	"testing"

	"github.com/JetBlink/orderbook/base"
	"github.com/shopspring/decimal"
)

func TestOrderBook_AddAsk(t *testing.T) {
	orderBook := NewOrderBook()
	order, err := NewOrder("abc", base.AskSide, "10", "20", uint64(100), nil)
	if err != nil {
		t.Error(err)
	}

	orderBook.AddOrder(order)
	orderBook.AddOrder(order)

	log.Println(base.ToJsonString(orderBook))
}

func TestOrderBook_AddAskOrder(t *testing.T) {
	orderBook := NewOrderBook()

	//价格第二
	order, err := NewOrder(
		"a2",
		base.AskSide,
		"1.102",
		"2.2",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//价格第一
	order, err = NewOrder(
		"a1",
		base.AskSide,
		"1.101",
		"1.1",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//价格第三
	order, err = NewOrder(
		"a3",
		base.AskSide,
		"1.103",
		"3.3",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//价格第二，时间同之前
	//order, err = NewOrder(
	//	"a4",
	//	"1.102",
	//	"4.4",
	//	"1559300877511654100",
	//)
	//if err != nil {
	//	t.Error(err)
	//}
	//orderBook.AddOrder(helper.AskSide, order)

	//价格第二，时间最后
	order, err = NewOrder(
		"a5",
		base.AskSide,
		"1.102",
		"5.5",
		1559300877511654101,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	if orderBook.Asks.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	if orderBook.Asks.Len() != 4 {
		t.Error("length error")
	}

	log.Println(base.ToJsonString(orderBook))
}

func TestFullOrderBook_AddAskOrderWithExistOrder(t *testing.T) {
	orderBook := NewOrderBook()

	//价格第二
	order, err := NewOrder(
		"a2",
		base.BidSide,
		"1.102",
		"2.2",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//订单 id 相同
	order, err = NewOrder(
		"a2",
		base.BidSide,
		"1.102",
		"1.1",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	if orderBook.Bids.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	if orderBook.Bids.Len() != 1 {
		t.Error("length error")
	}

	log.Println(base.ToJsonString(orderBook))
}

func TestFullOrderBook_AddBidOrder(t *testing.T) {
	orderBook := NewOrderBook()

	//价格第二
	order, err := NewOrder(
		"a2",
		base.BidSide,
		"1.102",
		"2.2",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//价格第一
	order, err = NewOrder(
		"a1",
		base.BidSide,
		"1.101",
		"1.1",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//价格第三
	order, err = NewOrder(
		"a3",
		base.BidSide,
		"1.103",
		"3.3",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//价格第二，时间同之前
	//order, err = NewOrder(
	//	"a4",
	//	"1.102",
	//	"4.4",
	//	1559300877511654100,
	//)
	//if err != nil {
	//	t.Error(err)
	//}
	//orderBook.AddOrder(helper.BidSide, order)

	//价格第二，时间最后
	order, err = NewOrder(
		"a5",
		base.BidSide,
		"1.102",
		"5.5",
		1559300877511654101,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	if orderBook.Bids.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	if orderBook.Bids.Len() != 4 {
		t.Error("length error")
	}

	log.Println(base.ToJsonString(orderBook))
}

func TestFullOrderBook_RemoveAskOrder(t *testing.T) {
	orderBook := NewOrderBook()

	//价格第二
	order, err := NewOrder(
		"a2",
		base.BidSide,
		"1.102",
		"2.2",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//订单 id 相同
	order, err = NewOrder(
		"a2",
		base.BidSide,
		"1.102",
		"1.1",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.removeOrder(order)

	if orderBook.Bids.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("length error")
	}

	log.Println(base.ToJsonString(orderBook))
}

func TestFullOrderBook_RemoveBidOrderByIdReal(t *testing.T) {
	orderBook := NewOrderBook()
	msg := [][3]string{
		{"5cf10b0acdaba439e5ad45cd", "8277.9", "0.01888553"},
		{"5cf10b0a054b463aafef3f0e", "8276", "0.0394"},
		{"5cf10b0b89fc846389110382", "8274.7", "0.40079"},
		{"5cf10af5cdaba43a1e7338bc", "8273.8", "0.02430003"}, //特殊
		{"5cf10b0b4c06872c13b9be4c", "8271.4", "0.00926951"},
		{"5cf10b0a054b463a72465902", "8270.3", "0.518085"},
		{"5cf10b0b89fc846350e843a2", "8270", "0.139868"},
		{"5cf10b0bc788c671ea12981d", "8268.6", "0.302001"},
	}

	for index, elem := range msg {
		order, err := NewOrder(elem[0], base.BidSide, elem[1], elem[2], uint64(index), nil)
		if err != nil {
			t.Error(err)
		}
		orderBook.AddOrder(order)
	}

	//PrintFullOrderBook("before", orderBook)

	order, err := NewOrder(
		"5cf10b0a0bad452af739ed1d",
		base.BidSide,
		"8268.3",
		"0.08",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	log.Println("add", base.ToJsonString(orderBook))
	//orderBook.RemoveByOrderId(helper.BidSide, "5cf10b0a0bad452af739ed1d")
	//PrintFullOrderBook("add", orderBook)

	//价格第二
	order, err = NewOrder(
		"5cf10b0b89fc846350e8437a",
		base.BidSide,
		"8268.3",
		"0.3141",
		1559300877511654101,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	log.Println("after", base.ToJsonString(orderBook))

	orderBook.RemoveByOrderId("5cf10b0a0bad452af739ed1d")

	log.Println("remove", base.ToJsonString(orderBook))

	if orderBook.Bids.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	//data, _ := json.Marshal(orderBook)
	//dump(data)
}

func TestFullOrderBook_RemoveAskOrderById(t *testing.T) {
	orderBook := NewOrderBook()

	//价格第二
	order, err := NewOrder(
		"a2",
		base.BidSide,
		"1.102",
		"2.2",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)
	if orderBook.Bids.Len() != len(orderBook.OrderPool) && len(orderBook.OrderPool) != 1 {
		t.Error("len error: 1")
	}

	orderBook.RemoveByOrderId(order.OrderId)

	if orderBook.Bids.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	if orderBook.Bids.Len() != 0 {
		t.Error("length error")
	}

	log.Println(base.ToJsonString(orderBook))
}

func TestFullOrderBook_RemoveBidNotExistOrder(t *testing.T) {
	orderBook := NewOrderBook()

	//价格第二
	order, err := NewOrder(
		"a2",
		base.BidSide,
		"1.102",
		"2.2",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//price和time相同
	order, err = NewOrder(
		"a5",
		base.BidSide,
		"1.102",
		"5.5",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.removeOrder(order)

	if orderBook.Bids.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	if orderBook.Bids.Len() != 1 {
		t.Error("length error")
	}

	log.Println(base.ToJsonString(orderBook))
}

func TestFullOrderBook_RemoveBidOrderOfAskOrder(t *testing.T) {
	orderBook := NewOrderBook()

	//价格第二
	order, err := NewOrder(
		"a2",
		base.BidSide,
		"1.102",
		"2.2",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	order, err = NewOrder(
		"a2",
		base.AskSide,
		"1.102",
		"2.2",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	//删除反方向订单
	orderBook.removeOrder(order)

	if orderBook.Bids.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	if orderBook.Bids.Len() != 1 {
		t.Error("length error")
	}

	log.Println(base.ToJsonString(orderBook))
}

func TestFullOrderBook_MatchOrder(t *testing.T) {
	orderBook := NewOrderBook()

	order, err := NewOrder(
		"a1",
		base.AskSide,
		"1.102",
		"1.1",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//价格 时间都相同
	order, err = NewOrder(
		"a2",
		base.AskSide,
		"1.102",
		"2.2",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	if err := orderBook.MatchOrder("a1", decimal.NewFromFloat(0.2)); err != nil {
		t.Error(err)
	}

	if !orderBook.GetOrder("a1").Size.Equal(decimal.NewFromFloat(0.9)) {
		t.Error("error size")
	}

	if orderBook.Asks.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	if orderBook.Asks.Len() != 2 {
		t.Error("length error")
	}

	log.Println(base.ToJsonString(orderBook))
}

func TestFullOrderBook_MatchOrderWithDone(t *testing.T) {
	orderBook := NewOrderBook()

	order, err := NewOrder(
		"a1",
		base.AskSide,
		"1.102",
		"1.1",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	if err := orderBook.MatchOrder(order.OrderId, order.Size); err != nil {
		t.Error(err)
	}

	if orderBook.Asks.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	if orderBook.Asks.Len() != 0 {
		t.Error("length error")
	}

	log.Println(base.ToJsonString(orderBook))
}

func TestFullOrderBook_ChangeOrder(t *testing.T) {
	orderBook := NewOrderBook()

	order, err := NewOrder(
		"a1",
		base.AskSide,
		"1.102",
		"1.1",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	//价格 时间都相同
	order, err = NewOrder(
		"a2",
		base.AskSide,
		"1.102",
		"2.2",
		1559300877511654100,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	orderBook.AddOrder(order)

	if err := orderBook.ChangeOrder("a1", decimal.NewFromFloat(0.2)); err != nil {
		t.Error(err)
	}

	if !orderBook.GetOrder("a1").Size.Equal(decimal.NewFromFloat(0.2)) {
		t.Error("error size")
	}

	if orderBook.Asks.Len() != len(orderBook.OrderPool) {
		t.Error("len error")
	}

	if orderBook.Asks.Len() != 2 {
		t.Error("length error")
	}

	log.Println(base.ToJsonString(orderBook))
}
