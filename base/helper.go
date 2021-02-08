package base

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// ToJsonString converts any value to JSON string.
func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

const (
	//ask 是 要价，喊价 卖家 卖单 Sort price from low to high
	AskSide = "asks"
	//bid 是 投标，买家 买单 Sort price from high to low
	BidSide = "bids"
)

func CheckSide(side string) error {
	switch side {
	case AskSide:
	case BidSide:
	default:
		return errors.New(fmt.Sprintf("error side, side should be %s or %s", AskSide, BidSide))
	}

	return nil
}
