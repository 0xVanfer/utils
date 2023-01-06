package utils

import (
	"errors"
	"math"
	"math/big"

	"github.com/0xVanfer/types"
)

// Curve stable coin model.
//
// Example:
//
//	info := &CurveStableCoinModel{
//		A:     50,
//		N:     2,
//		X_Num: 0,
//		Y_Num: 1,
//		Fee:   0.0004,
//		Balances: []*big.Int{
//			big.NewInt(1).Mul(big.NewInt(409880), big.NewInt(1e18)),
//			big.NewInt(1).Mul(big.NewInt(244647), big.NewInt(1e18)),
//		},
//		Dx: big.NewInt(1).Mul(big.NewInt(1), big.NewInt(1e18)),
//	}
//	info.CalcDy()
//	utils.PrettyJsonPrintln(info)
type CurveStableCoinModel struct {
	// Input info.
	A   int64   // Amplification. A != 0.
	N   int64   // Token number. N >= 2.
	Fee float64 // Curve Fee. Usually 0.0004.

	// Input info.
	// Token amount info. Must be multiplied by 1e18.
	Balances []*big.Int // Pool balance of tokens. len(Balances) == N.
	Dx       *big.Int   // Swap amount of token x.
	X_Num    int64      // The serial number of the token to be swaped. Related to `Balances`.
	Y_Num    int64      // The serial number of the token to be swaped to. Related to `Balances`.

	// The data for the convenience of calculation.
	sum  *big.Int // Sum of all balances.
	prod *big.Int // Production of all balances.
	an   *big.Int // A * n
	ann  *big.Int // A * n**n

	// Output info.
	D           *big.Int // Variable D of the pool.
	ExpectedFee *big.Int // Expected fee to be charged.
	ExpectedDy  *big.Int // Expected give out amount of token y.
}

// Regulatory check of details.
//
// Requirements:
//
//	A != 0;
//	len(Balances) == N >= 2;
//	Any of Balances !=0;
//	X_Num != Y_Num;
//	0 <= X_Num, Y_Num < N;
func (details *CurveStableCoinModel) regulatoryCheck() (err error) {
	if details.A == 0 {
		//lint:ignore ST1005 sentence error needed
		err = errors.New("Require: A != 0.")
		return
	}
	if details.N < 2 {
		//lint:ignore ST1005 sentence error needed
		err = errors.New("Require: N >= 2.")
		return
	}
	if int64(len(details.Balances)) != details.N {
		//lint:ignore ST1005 sentence error needed
		err = errors.New("Require: len(Balances) == N.")
		return
	}
	for _, balance := range details.Balances {
		if balance.Cmp(big.NewInt(0)) == 0 {
			//lint:ignore ST1005 sentence error needed
			err = errors.New("Require: No 0 in `Balances`.")
			return
		}
	}
	if details.X_Num == details.Y_Num {
		//lint:ignore ST1005 sentence error needed
		err = errors.New("Require: X_Num != Y_Num.")
		return
	}
	if details.X_Num < 0 || details.X_Num >= details.N {
		//lint:ignore ST1005 sentence error needed
		err = errors.New("Require: 0 <= X_Num < N.")
		return
	}
	if details.Y_Num < 0 || details.Y_Num >= details.N {
		//lint:ignore ST1005 sentence error needed
		err = errors.New("Require: 0 <= Y_Num < N.")
		return
	}
	if details.Dx == nil {
		//lint:ignore ST1005 sentence error needed
		err = errors.New("Require: Dx != 0.")
		return
	}
	return
}

// Calculate sum, prod, an, ann.
func (details *CurveStableCoinModel) calcSP() {
	details.an = big.NewInt(details.A * details.N)
	details.ann = big.NewInt(1).Mul(big.NewInt(1).Exp(big.NewInt(details.N), big.NewInt(details.N), nil), big.NewInt(details.A))
	// fmt.Println(details.ann)

	newSum := big.NewInt(0)
	newProd := big.NewInt(1)
	for _, balance := range details.Balances {
		// sum += balance
		newSum = big.NewInt(1).Add(newSum, balance)
		// prod += balance
		newProd = big.NewInt(1).Mul(newProd, balance)
	}
	details.sum = newSum
	details.prod = newProd

}

// Must be the first to be called.
func (details *CurveStableCoinModel) calcD() {
	details.calcSP()
	// D of the previous round.
	DPrev := big.NewInt(0)
	_ = DPrev
	// D is initialized as details.sum
	D := details.sum

	for i := 0; i < 255; i++ {
		D_P := D
		for _, balance := range details.Balances {
			// dp = dp * D / (balance * N)
			D_P = big.NewInt(1).Div(big.NewInt(1).Mul(D_P, D), big.NewInt(1).Mul(balance, big.NewInt(details.N)))
		}
		DPrev = D
		// (An * sum + D_P * N) * D
		up_ := big.NewInt(1).Mul(D, big.NewInt(1).Add(big.NewInt(1).Mul(details.an, details.sum), big.NewInt(1).Mul(D_P, big.NewInt(details.N))))
		// (An - 1) * D + (N + 1) * D_P
		down_ := big.NewInt(1).Add(big.NewInt(1).Mul(big.NewInt(1).Sub(details.an, big.NewInt(1)), D), big.NewInt(1).Mul(big.NewInt(details.N+1), D_P))
		D = big.NewInt(1).Div(up_, down_)
		// Abs of D - DPrev.
		diff := big.NewInt(1).Abs(big.NewInt(1).Sub(D, DPrev))
		// diff <= 1 return D
		if diff.Cmp(big.NewInt(1)) != 1 {
			details.D = D
			return
		}
	}
}

// Calculate everything.
func (details *CurveStableCoinModel) CalcDy() (err error) {
	err = details.regulatoryCheck()
	if err != nil {
		return
	}
	// Calculate the previous D.
	details.calcD()

	var an *big.Int = big.NewInt(details.A * details.N)
	var sum *big.Int = big.NewInt(0)
	var prod *big.Int = big.NewInt(1)
	var D_P *big.Int = details.D
	var yStart *big.Int

	for i := 0; i < int(details.N); i++ {
		var balance *big.Int
		if i == int(details.X_Num) {
			// If is the token to be swaped, continue, balance += dx.
			balance = big.NewInt(1).Add(details.Balances[i], details.Dx)
		} else if i == int(details.Y_Num) {
			yStart = details.Balances[i]
			// If is the token to be swaped to, continue.
			// Neither sum nor product will use this.
			continue
		} else {
			balance = details.Balances[i]
		}
		sum = big.NewInt(1).Add(sum, balance)
		prod = big.NewInt(1).Mul(prod, balance)
		D_P = big.NewInt(1).Div(big.NewInt(1).Mul(D_P, details.D), big.NewInt(1).Mul(balance, big.NewInt(details.N)))
	}

	c := big.NewInt(1).Div(big.NewInt(1).Mul(D_P, details.D), big.NewInt(1).Mul(an, big.NewInt(details.N)))
	b := big.NewInt(1).Add(sum, big.NewInt(1).Div(details.D, an))
	y := details.D
	yPrev := big.NewInt(0)
	_ = yPrev

	for i := 0; i < 255; i++ {
		yPrev = y
		// y * y +c
		up_ := big.NewInt(1).Add(big.NewInt(1).Mul(y, y), c)
		// 2 * y + b - D
		down_ := big.NewInt(1).Sub(big.NewInt(1).Add(big.NewInt(1).Mul(big.NewInt(2), y), b), details.D)
		y = big.NewInt(1).Div(up_, down_)

		// Abs of y - yPrev.
		diff := big.NewInt(1).Abs(big.NewInt(1).Sub(y, yPrev))
		// diff <= 1, y is decided
		if diff.Cmp(big.NewInt(1)) != 1 {
			dy := big.NewInt(1).Sub(yStart, y)
			details.ExpectedFee = big.NewInt(1).Div(big.NewInt(1).Mul(dy, big.NewInt(types.ToInt64(details.Fee*math.Pow10(18)))), big.NewInt(int64(math.Pow10(18))))
			details.ExpectedDy = big.NewInt(1).Sub(dy, details.ExpectedFee)
			// fmt.Println(details.ExpectedFee)
			// fmt.Println(details.ExpectedDy)
			return
		}
	}
	return
}
