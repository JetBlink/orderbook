package helper

import "encoding/json"

//func DecimalDiff(a string, b string) (bool, error) {
//	if a == b {
//		return true, nil
//	}
//
//	aF, err := decimal.NewFromString(a)
//	if err != nil {
//		return false, err
//	}
//	bF, err := decimal.NewFromString(b)
//	if err != nil {
//		return false, err
//	}
//
//	if !aF.Equal(bF) {
//		return false, nil
//	}
//
//	return true, nil
//}
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
