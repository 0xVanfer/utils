package utils

import (
	"fmt"
	"math/big"

	"github.com/0xVanfer/types"
)

// Curve stable token swap model.
// Two tokens vesion.
type CurveStableModel struct {
	// Input info.
	A int64 // Amplification.

	// Input info.
	// Token amount info. Must be multiplied by 1e18.
	BalanceX *big.Int // Pool balance of token x.
	BalanceY *big.Int // Pool balance of token y.
	Dx       *big.Int // Swap amount of token x.

	// Output info.
	D          *big.Int // Variable D of the pool.
	ExpectedDy *big.Int // Expected give out amount of token y.
	Slippage   float64  // The slippage
}

// Calculate D of the pool, expected dy and slippage of the swap.
// A, BalanceX, BalanceY must not be nil.
func (details *CurveStableModel) CalcDy() {
	if details.A == 0 {
		fmt.Println("A cannot be 0")
		return
	}
	if details.BalanceX == nil {
		fmt.Println("BalanceX cannot be nil")
		return
	}
	if details.BalanceY == nil {
		fmt.Println("BalanceY cannot be nil")
		return
	}
	x1 := details.BalanceX
	x2 := details.BalanceY
	// balance1 + balance2; x1 add x2
	x1ax2 := big.NewInt(1).Add(x1, x2)
	// balance1 * balance2; x1 mul x2
	x1mx2 := big.NewInt(1).Mul(x1, x2)

	// https://docs.balancer.fi/concepts/math/stable-math
	// To get D, solve the function:
	// D^3 + (16 * A * X1 * X2 - 4 * X1 * X2) * D - 16 * A * X1 * X2 * (X1 + X2) = 0
	// D^3 + p * D + q = 0
	// p = 16 * A * X1 * X2 - 4 * X1 * X2
	p := big.NewInt(1).Sub(big.NewInt(1).Mul(big.NewInt(16*details.A), x1mx2), big.NewInt(1).Mul(big.NewInt(4), x1mx2))
	// q = - 16 * A * X1 * X2 * (X1 + X2)
	q := big.NewInt(1).Mul(big.NewInt(-16*details.A), big.NewInt(1).Mul(x1ax2, x1mx2))

	// p/3
	pd3 := big.NewInt(1).Div(p, big.NewInt(3))
	// q/2
	qd2 := big.NewInt(1).Div(q, big.NewInt(2))

	// (q/2)^2
	i2q := big.NewInt(1).Exp(qd2, big.NewInt(2), nil)
	// (p/3)^3
	i3p := big.NewInt(1).Exp(pd3, big.NewInt(3), nil)
	// (i2q + i3p)^1/2
	internal := BigSqrt(big.NewInt(1).Add(i2q, i3p), 2)

	// (-q/2 + internal)^1/3
	re1 := BigSqrt(big.NewInt(1).Sub(internal, qd2), 3)
	// (-q/2 - internal)^1/3 = - (q/2 + internal)^1/3
	re2 := big.NewInt(1).Mul(BigSqrt(big.NewInt(1).Add(internal, qd2), 3), big.NewInt(-1))

	// D = re1 + re2
	D := big.NewInt(1).Add(re1, re2)
	details.D = D

	// Do not calculate the dy.
	dx1 := details.Dx
	if (dx1 == nil) || (dx1.Cmp(big.NewInt(0)) == 0) {
		return
	}
	// x1 + dx1
	x1adx1 := big.NewInt(1).Add(x1, dx1)

	// Trading.
	// https://dev.balancer.fi/resources/pool-math/stable-math
	// Swap token1 for token2.
	// y^2 + (D/4A + (X1 + dX1) - D) * y - D^3 / (16A * (X1 + dX1)) = 0
	// y^2 + b * x + c = 0
	// b = D/4A + (X1 + dX1) - D
	b := big.NewInt(1).Sub(big.NewInt(1).Add(big.NewInt(1).Div(D, big.NewInt(4*details.A)), x1adx1), D)
	// c = -D^3 / (16A * (X1 + dX1))
	c := big.NewInt(1).Mul(big.NewInt(1).Div(big.NewInt(1).Exp(D, big.NewInt(3), nil), big.NewInt(1).Mul(big.NewInt(16*details.A), x1adx1)), big.NewInt(-1))

	// (b^2 - 4 * a * c)^1/2; a = 1
	delta := BigSqrt(big.NewInt(1).Sub(big.NewInt(1).Exp(b, big.NewInt(2), nil), big.NewInt(1).Mul(big.NewInt(4), c)), 2)
	// y = (-b + delta) / 2
	y := big.NewInt(1).Div(big.NewInt(1).Sub(delta, b), big.NewInt(2))
	// dx2 = x2 - y
	dx2 := big.NewInt(1).Sub(x2, y)
	details.ExpectedDy = dx2
	// Price difference of dx1 and dx2.
	diff := big.NewInt(1).Abs(big.NewInt(1).Sub(dx1, dx2))
	slippage := types.ToFloat64(diff) / types.ToFloat64(big.NewInt(1).Abs(details.Dx))
	details.Slippage = slippage
}
