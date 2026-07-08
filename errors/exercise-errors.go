package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return 0, e
	}
	
	z := 1.0

	for i := 0; i < 50; i++ {
		oldZ := z
		z -= (z*z - x) / (2 * z)

		if math.Abs(z-oldZ) < 1e-9 {
			return z, nil
		}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(-2))
}
