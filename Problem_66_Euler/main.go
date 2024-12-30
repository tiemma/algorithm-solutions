package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
	"time"
)

func getK(a, b, D *big.Float) *big.Float {
	aSquared := new(big.Float).Mul(a, a)
	bSquared := new(big.Float).Mul(b, b)

	return new(big.Float).Sub(aSquared, new(big.Float).Mul(D, bSquared))
}

func composeTriples(a, b, D, k, m *big.Float) (*big.Float, *big.Float, *big.Float) {
	absK := new(big.Float).Abs(k)

	am := new(big.Float).Mul(a, m)
	db := new(big.Float).Mul(D, b)
	amdb := new(big.Float).Add(am, db)

	bm := new(big.Float).Mul(b, m)
	abm := new(big.Float).Add(a, bm)

	m2 := new(big.Float).Mul(m, m)
	m2D := new(big.Float).Sub(m2, D)

	return new(big.Float).Quo(amdb, absK), new(big.Float).Quo(abm, absK), new(big.Float).Quo(m2D, k)
}

// BrahmaguptaConsumption
// Apply https://en.wikipedia.org/wiki/Chakravala_method#Brahmagupta's_composition_method to derive solutions for known versions of k
type BrahmaguptaConsumption func(a, b, k *big.Float) (*big.Float, *big.Float)

func k1(a, b, k *big.Float) (*big.Float, *big.Float) {
	a2 := new(big.Float).Mul(a, a)
	doubleA2 := new(big.Float).Mul(big.NewFloat(2), a2)

	x := new(big.Float).Add(doubleA2, big.NewFloat(float64(-1*k.Sign())))

	ab := new(big.Float).Mul(a, b)
	y := new(big.Float).Mul(ab, big.NewFloat(2))
	return x, y
}

func k2(a, b, k *big.Float) (*big.Float, *big.Float) {
	a2 := new(big.Float).Mul(a, a)
	x := a2.Add(a2, big.NewFloat(float64(-1*k.Sign())))

	y := new(big.Float).Mul(a, b)
	return x, y
}

func k4(a, b, k *big.Float) (*big.Float, *big.Float) {
	one := big.NewFloat(1)
	two := big.NewFloat(2)
	three := big.NewFloat(3)
	if k.Sign() > 0 {
		if new(big.Float).Quo(k, b).IsInt() {
			a2 := new(big.Float).Mul(a, a)
			a2m2 := new(big.Float).Sub(a2, two)

			x := new(big.Float).Quo(a2m2, two)

			ab := new(big.Float).Mul(a, b)
			y := new(big.Float).Quo(ab, two)

			return x, y
		} else {
			a2 := new(big.Float).Mul(a, a)
			a2m3 := new(big.Float).Sub(a2, three)

			ad2 := new(big.Float).Quo(a, two)
			x := new(big.Float).Mul(ad2, a2m3)

			a2m1 := new(big.Float).Sub(a2, one)
			bd2 := new(big.Float).Quo(b, two)
			y := new(big.Float).Mul(bd2, a2m1)

			return x, y
		}
	}

	a2 := new(big.Float).Mul(a, a)
	a2p1 := new(big.Float).Add(a2, one)
	a2p2 := new(big.Float).Add(a2, two)
	a2p3 := new(big.Float).Add(a2, three)

	a2p1ma2p3 := new(big.Float).Mul(a2p1, a2p3)
	a2p1ma2p3m2 := new(big.Float).Sub(a2p1ma2p3, two)
	a2p1ma2p3m2ma2p2 := new(big.Float).Mul(a2p1ma2p3m2, a2p2)
	x := new(big.Float).Quo(a2p1ma2p3m2ma2p2, two)

	ab := new(big.Float).Mul(a, b)
	a2p1ma2p3mab := new(big.Float).Mul(a2p1ma2p3, ab)
	y := new(big.Float).Quo(a2p1ma2p3mab, two)

	//Note, k=-4 is useful to find a solution to Pell's Equation, but it is not always the smallest integer pair.
	//panic("Do not apply when k=-4")
	return x, y
}

func findMinimalM(a, b, D, k *big.Float) *big.Float {
	inflectionPoint := big.NewFloat(math.MaxInt64)
	lastM := new(big.Float)
	lastK := big.NewFloat(math.MaxInt64)
	m := big.NewFloat(1)
	incrementer := big.NewFloat(1)

	value, _ := k.Float64()
	if value < 1 {
		value = 1 * float64(k.Sign())
	}
	kShift := big.NewFloat(float64(int(value) * k.Sign()))
	for {
		if m.Cmp(big.NewFloat(1000)) == 1 {
			return lastM
		}

		bigM := new(big.Float).Copy(m)
		_, newB, newK := composeTriples(a, b, D, k, bigM)

		absNewK := new(big.Float).Abs(newK)
		absNewK.Mul(absNewK, k)
		absNewK = new(big.Float).Abs(absNewK)

		if newB.IsInt() && absNewK.IsInt() {
			if absNewK.Cmp(inflectionPoint) == 1 && newB.Cmp(new(big.Float)) == 1 {
				return lastM
			}

			inflectionPoint = absNewK
			if lastK.Cmp(absNewK) == 1 {
				lastK = absNewK
				lastM = bigM
			}

			incrementer = kShift
		}

		m.Add(m, incrementer)
	}
}

func isExemptK(k *big.Float) bool {
	kAbs := new(big.Float).Abs(k)
	return kAbs.Cmp(big.NewFloat(1)) == 0 || kAbs.Cmp(big.NewFloat(2)) == 0 || k.Cmp(big.NewFloat(4)) == 0
}

func PellsEquation(window float64) {

	solutions := map[float64]BrahmaguptaConsumption{
		1: k1,
		2: k2,
		4: k4,
	}
	maxIterations := 200
	skipped := 0

	// Pell's Equation
	// https://en.wikipedia.org/wiki/Pell%27s_equation#Solutions
	// x^2 - D * y^2 = 1
	// a^2 - D * b^2 = k

	// Uses the Chakravala method: http://en.wikipedia.org/wiki/Chakravala_method
	maxD, x, y := big.NewFloat(0), big.NewFloat(math.MinInt64), big.NewFloat(0)
	for i := 2.0; i <= window; i++ {
		D := big.NewFloat(i)
		sqrtD := new(big.Float).Sqrt(D)
		if sqrtD.IsInt() {
			continue
		}

		floatSqrtD := new(big.Float).Copy(sqrtD).SetPrec(1)
		b := big.NewFloat(1)
		a, k := floatSqrtD, getK(floatSqrtD, b, D)
		kp1 := new(big.Float).Add(floatSqrtD, big.NewFloat(1))
		newK := getK(kp1, b, D)
		isSqrt := true

		if new(big.Float).Abs(newK).Cmp(new(big.Float).Abs(k)) == -1 {
			a, k = kp1, newK
		}

		fmt.Println("--------------------")
		fmt.Printf("Starting %f, %f, %f, %f\n", D, a, b, k)
		originalA, originalB, originalK := new(big.Float).Copy(a), new(big.Float).Copy(b), new(big.Float).Copy(k)

		iterations := 0
		disableSqrt := false
		for {
			if !isSqrt && isExemptK(k) {
				break
			}

			m := findMinimalM(a, b, D, k)
			if m.Cmp(new(big.Float)) < 0 {
				panic("m cannot be less than 0")
			}
			fmt.Printf("Before %f, %f, %f, %f\n", a, b, k, m)

			a, b, k = composeTriples(a, b, D, k, m)
			isSqrt = false

			kAbs := new(big.Float).Abs(k)
			kSqrt := new(big.Float).Sqrt(kAbs)
			if kSqrt.IsInt() && !isExemptK(k) && !disableSqrt {
				fmt.Println("kSquared Before", a, b, k)
				isSqrt = true
				a.Quo(a, kSqrt)

				b.Quo(b, kSqrt)
				k.Quo(k, kAbs)
			}

			if iterations%maxIterations == 0 {
				if disableSqrt {
					break
				}

				fmt.Println("Disabling sqrt")
				a = new(big.Float).Copy(originalA)
				b = new(big.Float).Copy(originalB)
				k = new(big.Float).Copy(originalK)
				disableSqrt = true
			}

			fmt.Printf("After %f, %f, %f\n", a, b, k)
			iterations++
		}

		index, _ := new(big.Float).Abs(k).Float64()
		if !k.IsInt() || index > 4 {
			skipped += 1
			continue
		}
		newA, newB := solutions[index](a, b, k)

		fmt.Printf("*****Solved in %d iterations**********\n", iterations)
		fmt.Println(k)
		fmt.Printf("Base solution %f, %f, %f\n", D, a, b)
		if k.Cmp(big.NewFloat(1)) == 0 {
			if a.Cmp(x) == 1 {
				x = a
				y = b
				maxD = D
			}
		} else {
			fmt.Printf("%f, %f, %f\n", D, newA, newB)

			if newA.Cmp(x) == 1 {
				x = newA
				y = newB
				maxD = D
			}
		}
	}

	fmt.Println("------------")
	fmt.Printf("Skippped %d operations\n", skipped)
	fmt.Printf("%f, %f, %f\n", maxD, x, y)
}

func getDenominator(number, lastRoot int) int {
	return number - int(math.Pow(float64(lastRoot), 2))
}

func solvePell(D float64) (float64, float64, int) {
	t := math.Sqrt(D)
	a := math.Floor(t)
	iterations := 0
	if a == t {
		return 0, 0, iterations
	}
	b, c, r := a, 1.0, float64(int64(a)<<1)
	e1, e2 := 1.0, 0.0
	f1, f2 := 0.0, 1.0
	for {
		b = (r * c) - b
		c = math.Floor((D - (b * b)) / c)
		r = math.Floor((a + b) / c)
		e1, e2 = e2, e1+e2*r
		f1, f2 = f2, f1+f2*r
		x, y := f2*a+e2, f2
		fmt.Printf("--%f %f\n", x, y)
		time.Sleep(500 * time.Millisecond)
		//sqrtD := big.NewFloat(t)

		// Go has some precision issue that makes this round off to 0 instead of 1
		// for a number of solutions.
		// It is highly impractical but in the future, just use the PellersEquation() method
		// as that works fine
		x2 := x + (t * y)
		Dy2 := x - (t * y)
		fmt.Printf("%f %f %f %f\n", x, y, x2, Dy2)

		if x2*Dy2 == 1 {
			return x, y, iterations
		}
		iterations += 1
	}
}

func getNumerator(D float64) (*big.Float, *big.Float, int) {
	x := math.Floor(math.Sqrt(D))
	y, z, r := x, 1.0, float64(int(x)<<1)
	x1, x2 := big.NewFloat(1), new(big.Float)
	y1, y2 := new(big.Float), big.NewFloat(1)
	iterations := 0
	for {
		y = r*z - y
		z = math.Floor((D - math.Pow(y, 2)) / z)
		r = math.Floor((x + y) / z)

		x1, x2 = x2, new(big.Float).Copy(x1).Add(x1, new(big.Float).Mul(x2, big.NewFloat(r)))
		y1, y2 = y2, new(big.Float).Copy(x1).Add(y1, new(big.Float).Mul(y2, big.NewFloat(r)))

		a, b := new(big.Float).Add(x2, new(big.Float).Mul(y2, big.NewFloat(x))), y2
		//if a.IsInf() || b.IsInf() || iterations > 200 {
		//	return new(big.Float), new(big.Float), iterations
		//}
		a2 := new(big.Float).Mul(a, a)
		b2 := new(big.Float).Mul(b, b)
		remainder := new(big.Float).Sub(a2, new(big.Float).Mul(big.NewFloat(D), b2))
		fmt.Println(remainder)
		time.Sleep(500 * time.Millisecond)
		if remainder.Cmp(big.NewFloat(1)) == 0 {
			return a, b, iterations
		}

		iterations += 1
	}
}

func main() {
	window := flag.Float64("window", 10, "window size")
	flag.Parse()

	PellsEquation(*window)

	//maxX := float64(math.MinInt64)
	//y := 0.0
	//D := 0.0
	//skipped := 0
	//
	//for i := 61.0; i <= *window; i++ {
	//	Dsqrt := math.Sqrt(i)
	//	if math.Floor(Dsqrt)-Dsqrt == 0 {
	//		continue
	//	}
	//	num, denum, iterations := solvePell(i)
	//	fmt.Println(i, num, denum, iterations)
	//	if iterations > 200 {
	//		skipped += 1
	//	}
	//	if num > maxX {
	//		maxX = num
	//		D = i
	//		y = denum
	//	}
	//}
	////time.Sleep(500 * time.Millisecond)
	//
	//fmt.Println(D, maxX, y, skipped)
}
