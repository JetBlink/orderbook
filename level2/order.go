package level2

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

//Price, Quantity
type Order struct {
	Price decimal.Decimal
	Size  decimal.Decimal
}

func NewOrder(price string, size string) (order *Order, err error) {
	priceValue, err := decimal.NewFromString(price)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("NewLevel3Order failed, price: `%s`, error: %v", price, err))
	}

	sizeValue, err := decimal.NewFromString(size)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("NewLevel3Order failed, size: `%s`, error: %v", size, err))
	}

	order = &Order{
		Price: priceValue,
		Size:  sizeValue,
	}

	return
}
