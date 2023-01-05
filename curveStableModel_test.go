package utils

import (
	"math/big"
	"testing"
)

func TestCurveStableCoinModel(t *testing.T) {
	big5k := big.NewInt(1).Mul(big.NewInt(5000), big.NewInt(1e18))
	_ = big5k
	big5m := big.NewInt(1).Mul(big.NewInt(5000000), big.NewInt(1e18))
	_ = big5m

	info := &CurveStableCoinModel{
		A:     50,
		N:     2,
		X_Num: 0,
		Y_Num: 1,
		Fee:   0.0004,
		Balances: []*big.Int{
			big.NewInt(1).Mul(big.NewInt(409880), big.NewInt(1e18)),
			big.NewInt(1).Mul(big.NewInt(244647), big.NewInt(1e18)),
		},
		Dx: big.NewInt(1).Mul(big.NewInt(1), big.NewInt(1e18)),
	}
	info.CalcDy()
	PrettyJsonPrintln(info)
	// fmt.Println(info.sum, info.ann, info.prod)

}
