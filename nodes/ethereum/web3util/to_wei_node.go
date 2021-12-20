package web3util

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func DecimalToWei(amount decimal.Decimal, decimals int) *big.Int {
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := big.NewInt(0)
	wei.SetString(result.String(), 10)
	return wei
}
