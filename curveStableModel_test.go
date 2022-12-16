package utils

import (
	"testing"

	"github.com/0xVanfer/types"
)

func TestCurveStableModel(t *testing.T) {
	info := &CurveStableModel{
		A:        1000,
		BalanceX: types.ToBigInt(1).Mul(types.ToBigInt(5000000), types.ToBigInt(1e18)),
		BalanceY: types.ToBigInt(1).Mul(types.ToBigInt(5000000), types.ToBigInt(1e18)),
		Dx:       types.ToBigInt(1).Mul(types.ToBigInt(-1000), types.ToBigInt(1e18)),
	}
	info.CalcDy()
	PrettyJsonPrintln(info)
}
