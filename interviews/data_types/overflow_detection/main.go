package main

import (
	"errors"
	"math"
)

var ErrIntOverflow = errors.New("integer overflow")

func Inc(counter int) (int, error) {
	if counter == math.MaxInt {
		return 0, ErrIntOverflow
	}

	return counter + 1, nil
}

func Add(lhs, rhs int) (int, error) {
	if rhs > 0 {
		if lhs > math.MaxInt-rhs {
			return 0, ErrIntOverflow
		}
	} else if rhs < 0 {
		if lhs < math.MinInt-rhs {
			return 0, ErrIntOverflow
		}
	}

	return lhs + rhs, nil
}

func Mul(lhs, rhs int) (int, error) {
	if lhs == 0 || rhs == 0 {
		return 0, nil
	}

	if lhs == 1 || rhs == 1 {
		return lhs * rhs, nil
	}

	if (lhs == -1 && rhs == math.MinInt) || (rhs == -1 && lhs == math.MinInt) {
		return 0, ErrIntOverflow
	}

	if lhs > math.MaxInt/rhs {
		return 0, ErrIntOverflow
	} else if lhs < math.MinInt/rhs {
		return 0, ErrIntOverflow
	}

	return lhs * rhs, nil
}
