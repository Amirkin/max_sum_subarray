package main

import (
	"math"
)

type segment struct {
	Sum   int
	Left  int
	Right int
}

func findSum(arr []int) *segment {
	var allNeg bool = true
	var negMax *segment = &segment{
		Sum:   math.MinInt32,
		Left:  0,
		Right: 0,
	}

	max := &segment{}
	current := &segment{}

	for i, el := range arr {

		current.Sum += el
		current.Right = i

		if el >= 0 {
			allNeg = false

			if current.Sum >= max.Sum {
				*max = *current
			}
		}

		if el < 0 {
			if allNeg {
				incNegCount(negMax, i, el)
			}

			if current.Sum <= 0 {
				current.Sum = 0
				current.Left = i + 1
			}
		}
	}

	if allNeg {
		return negMax
	}

	return max
}

func incNegCount(negMax *segment, i, el int) {
	if el > negMax.Sum {
		negMax.Sum = el
		negMax.Left = i
		negMax.Right = i
	}
}
