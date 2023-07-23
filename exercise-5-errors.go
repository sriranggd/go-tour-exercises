package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number : %f", float64(e))
}

func Sqrtv3(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return x, err
	}

	return Sqrtv2(x), nil
}
