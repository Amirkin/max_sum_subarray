package main

import (
	"math"
)

type Segment struct {
	Sum   int
	Left  int
	Right int
}

func findSum(arr []int) *Segment {
	var allNeg bool = true
	var negMax *Segment = &Segment{
		Sum:   math.MinInt32,
		Left:  0,
		Right: 0,
	}

	max := &Segment{}
	current := &Segment{}

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

func incNegCount(negMax *Segment, i, el int) {
	if el > negMax.Sum {
		negMax.Sum = el
		negMax.Left = i
		negMax.Right = i
	}
}
