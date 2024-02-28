package copypasta

import (
	"math"
	"math/big"
)

/*
https://cp-algorithms.com/algebra/continued-fractions.html
连分数中的常数 https://mathworld.wolfram.com/ContinuedFractionConstants.html
https://mathworld.wolfram.com/RamanujanContinuedFractions.html
https://mathworld.wolfram.com/RegularContinuedFraction.html
https://mathworld.wolfram.com/GeneralizedContinuedFraction.html
https://mathworld.wolfram.com/ModifiedBesselFunctionoftheFirstKind.html
https://mathworld.wolfram.com/ModifiedBesselFunctionoftheSecondKind.html
https://oeis.org/A001053
https://oeis.org/A001040
https://oeis.org/A052119
*/
func continuedFractionCollections() {
	// a/b = [exp[0]; exp[1],...]
	calcContinuedFraction := func(a, b int) (exp []int) {
		for b != 1 {
			exp = append(exp, a/b)
			a, b = b, a%b
		}
		exp = append(exp, a)
		return
	}
	calcContinuedFractionBig := func(a_, b_ int) (exp []*big.Int) {
		a, b := big.NewInt(int64(a_)), big.NewInt(int64(b_))
		for !b.IsInt64() || b.Int64() != 1 {
			r := &big.Int{}
			a.QuoRem(a, b, r)
			exp = append(exp, new(big.Int).Set(a))
			a, b = b, r
		}
		exp = append(exp, a)
		return
	}

	// https://codeforces.com/contest/281/problem/B
	// todo 见 Python fractions.Fraction.limit_denominator 源码

	// sqrt(d) = [exp[0]; exp[1],..., 2*exp[0], exp[1], ..., 2*exp[0], exp[1], ...]
	// https://en.wikipedia.org/wiki/Pell%27s_equation 解 https://oeis.org/A002350 https://oeis.org/A002349
	// https://www.weiwen.io/post/about-the-pell-equations-2/
	// 连分数表示 https://oeis.org/A240071
	// 循环节长度 https://oeis.org/A003285
	calcSqrtContinuedFraction := func(d int) (exp []int) {
		sqrtD := math.Sqrt(float64(d))
		base := int(sqrtD)
		if base*base == d {
			return []int{base}
		}
		p := []int{0}
		q := []int{1}
		for i := 0; ; i++ {
			a := int((float64(p[i]) + sqrtD) / float64(q[i]))
			exp = append(exp, a)
			if a == 2*base {
				break
			}
			pi := a*q[i] - p[i]
			p = append(p, pi)
			q = append(q, (d-pi*pi)/q[i]) // q[i] 可以整除 d-pi*pi
		}
		return
	}
	//calcSqrtContinuedFraction := func(d int) (exp []int) {
	//	sqrtD := math.Sqrt(float64(d))
	//	base0 := int(sqrtD)
	//	if base0*base0 == d {
	//		return []int{base0}
	//	}
	//	a := []int{1}
	//	b := []int{0}
	//	c := []int{1}
	//	for i := 0; ; i++ {
	//		base := int((float64(a[i])*sqrtD + float64(b[i])) / float64(c[i]))
	//		exp = append(exp, base)
	//		if base == 2*base0 {
	//			break
	//		}
	//		tmp := base*c[i] - b[i]
	//		newA := c[i] * a[i]
	//		newB := c[i] * tmp
	//		newC := d*a[i]*a[i] - tmp*tmp
	//		gcd := calcGCDN(newA, newB, newC)
	//		a = append(a, newA/gcd)
	//		b = append(b, newB/gcd)
	//		c = append(c, newC/gcd)
	//	}
	//	return
	//}

	gcd := func(a, b int) int {
		for b > 0 {
			a, b = b, a%b
		}
		return a
	}
	gcds := func(a ...int) (g int) {
		g = a[0]
		for _, v := range a[1:] {
			g = gcd(g, v)
		}
		return
	}
	// sqrt(m/n)
	calcSqrtRatContinuedFraction := func(m, n int) (exp []int) {
		sqrtRat := math.Sqrt(float64(m) / float64(n))
		base0 := int(sqrtRat)
		if base0*base0*n == m {
			return []int{base0}
		}
		a := []int{1}
		b := []int{0}
		c := []int{1}
		const loop = 50
		for i := 0; i < loop; i++ {
			base := int((float64(a[i])*sqrtRat + float64(b[i])) / float64(c[i]))
			exp = append(exp, base)
			tmp := base*c[i] - b[i]
			newA := n * c[i] * a[i]
			newB := n * c[i] * tmp
			newC := m*a[i]*a[i] - n*tmp*tmp
			if newC == 0 {
				return
			}
			g := gcds(newA, newB, newC)
			//Println(i, base, newA, newB, newC, g)
			a = append(a, newA/g)
			b = append(b, newB/g)
			c = append(c, newC/g)
		}
		return
	}
	calcSqrtRatContinuedFractionBig := func(m, n int) (exp []*big.Int) {
		sqrtRat := new(big.Float).Sqrt(big.NewFloat(float64(m) / float64(n)))
		if base0, acc := sqrtRat.Int(nil); acc == big.Exact {
			exp = append(exp, base0)
			return
		}
		bigM, bigN := big.NewInt(int64(m)), big.NewInt(int64(n))
		a := []*big.Int{big.NewInt(1)}
		b := []*big.Int{big.NewInt(0)}
		c := []*big.Int{big.NewInt(1)}
		const loop = 50
		for i := 0; i < loop; i++ {
			tmpF := new(big.Float).SetInt(a[i])
			base, _ := tmpF.Mul(tmpF, sqrtRat).
				Add(tmpF, new(big.Float).SetInt(b[i])).
				Quo(tmpF, new(big.Float).SetInt(c[i])).
				Int(nil)
			exp = append(exp, new(big.Int).Set(base))
			tmp := &big.Int{}
			tmp.Mul(base, c[i]).Sub(tmp, b[i])
			newA, newB, newC := &big.Int{}, &big.Int{}, &big.Int{}
			newA.Mul(bigN, c[i]).Mul(newA, a[i])
			newB.Mul(bigN, c[i]).Mul(newB, tmp)
			tmp.Mul(tmp, tmp).Mul(tmp, bigN)
			newC.Mul(bigM, a[i]).Mul(newC, a[i]).Sub(newC, tmp)
			if newC.IsInt64() && newC.Int64() == 0 {
				return
			}
			g := big.NewInt(1)
			if !base.IsInt64() || base.Int64() > 0 {
				g.GCD(nil, nil, newA, newB).GCD(nil, nil, g, newC)
			}
			//Println(i, base, newA, newB, newC, gcd)
			a = append(a, newA.Quo(newA, g))
			b = append(b, newB.Quo(newB, g))
			c = append(c, newC.Quo(newC, g))
		}
		return
	}

	// 将连分数化成最简分数
	// 模板题 https://leetcode-cn.com/contest/season/2019-fall/problems/deep-dark-fraction/
	calcRatByContinuedFraction := func(exp []int) (a, b int) {
		n := len(exp)
		h := make([]int, n+2)
		h[0], h[1] = 0, 1
		k := make([]int, n+2)
		k[0], k[1] = 1, 0
		for i, v := range exp {
			h[i+2] = v*h[i+1] + h[i]
			k[i+2] = v*k[i+1] + k[i]
		}
		return h[n+1], k[n+1]
	}

	_ = []any{
		calcContinuedFraction, calcContinuedFractionBig,
		calcSqrtContinuedFraction, calcSqrtRatContinuedFraction, calcSqrtRatContinuedFractionBig,
		calcRatByContinuedFraction,
	}
}
