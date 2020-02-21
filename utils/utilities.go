package utils

import (
	"fmt"
	"math/big"
)

const (
	Wei = 1e18
)

// ConvertFloatAmountToBigInt - converts a given float64 amount to a bigint with the correct base
func ConvertFloatAmountToBigInt(amount float64) *big.Int {
	bigAmount := new(big.Float).SetFloat64(amount)
	base := new(big.Float).SetInt(big.NewInt(Wei))
	bigAmount.Mul(bigAmount, base)
	realAmount := new(big.Int)
	bigAmount.Int(realAmount)

	return realAmount
}

// ConvertNumeralStringToBigInt - converts a numeral string back to a big float with the correct base set
func ConvertNumeralStringToBigInt(balance string) (*big.Int, error) {
	value := new(big.Float)
	value, ok := value.SetString(balance)

	if !ok {
		return nil, fmt.Errorf("can't convert balance string %s to a float balance", balance)
	}

	base := new(big.Float).SetInt(big.NewInt(Wei))
	value.Mul(value, base)

	result := new(big.Int)
	uintval, _ := value.Uint64()
	result.SetUint64(uintval)

	return result, nil
}
