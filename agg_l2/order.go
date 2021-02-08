package agg_l2

import (
	"fmt"

	"github.com/shopspring/decimal"
)

//聚合后的订单
//AggregatedPrice, Quantity
type aggregatedOrder struct {
	AggregatedPrice decimal.Decimal
	Size            decimal.Decimal
}

//原始订单
//Price, Quantity
type Order struct {
	Price decimal.Decimal
	Size  decimal.Decimal
	Info  interface{}
}

func NewOrder(price string, size string, info interface{}) (order *Order, err error) {
	priceValue, err := decimal.NewFromString(price)
	if err != nil {
		return nil, fmt.Errorf("NewOrder failed, price: `%s`, error: %v", price, err)
	}

	sizeValue, err := decimal.NewFromString(size)
	if err != nil {
		return nil, fmt.Errorf("NewOrder failed, size: `%s`, error: %v", size, err)
	}

	order = &Order{
		Price: priceValue,
		Size:  sizeValue,
		Info:  info,
	}

	return
}
