package utils

import (
	"math/big"
	"testing"
)

func TestCurveStableModel(t *testing.T) {
	info := &CurveStableModel{
		A:        1000,
		BalanceX: big.NewInt(1).Mul(big.NewInt(5000000), big.NewInt(1e18)),
		BalanceY: big.NewInt(1).Mul(big.NewInt(5000000), big.NewInt(1e18)),
		Dx:       big.NewInt(1).Mul(big.NewInt(-1000), big.NewInt(1e18)),
	}
	info.CalcDy()
	PrettyJsonPrintln(info)
}
