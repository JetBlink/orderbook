package level3

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

type Order struct {
	OrderId string
	Price   decimal.Decimal
	Size    decimal.Decimal
	Time    uint64
}

func NewOrder(orderId string, price string, size string, time uint64) (order *Order, err error) {
	priceValue, err := decimal.NewFromString(price)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("NewOrder failed, price: `%s`, error: %v", price, err))
	}

	sizeValue, err := decimal.NewFromString(size)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("NewOrder failed, size: `%s`, error: %v", size, err))
	}

	order = &Order{
		OrderId: orderId,
		Price:   priceValue,
		Size:    sizeValue,
		Time:    time,
	}
	return
}
