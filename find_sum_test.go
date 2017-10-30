package main

import (
	"testing"
)

func TestAllPositive(t *testing.T) {
	arr := []int{5, 6, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6, 3, 2, 3, 4, 5, 6, 6}

	var expectedSum int
	for _, item := range arr {
		expectedSum += item
	}

	checkSegment(arr, 0, len(arr)-1, expectedSum, t)
}

func TestAllNegative(t *testing.T) {
	arr := []int{-5, -5, -7, -1, -9, -3, -3, -4, -5}
	checkSegment(arr, 3, 3, -1, t)
}

func TestNegativeWithZero(t *testing.T) {
	arr := []int{-5, -5, -7, -1, 0, -3, -3, -4, -5}
	checkSegment(arr, 4, 4, 0, t)
}

func TestOneBiggestSegment(t *testing.T) {
	arr := []int{1, 2, 3, -6, 100, -1, -2, 3}
	checkSegment(arr, 4, 7, 100, t)
}

func TestBiggestSegmentWithNegative(t *testing.T) {
	arr := []int{1, 2, 3, -3, 5, -100, 7}
	checkSegment(arr, 0, 4, 8, t)
}

func TestStartNegative(t *testing.T) {
	arr := []int{-5, 1, 2, 3}
	checkSegment(arr, 1, 3, 6, t)
}

func TestTwoBiggestSegment(t *testing.T) {
	arr := []int{-5, 7, -4, 7}
	checkSegment(arr, 1, 3, 10, t)
}

func checkSegment(arr []int, left, right, sum int, t *testing.T) {
	seg := findSum(arr)

	if seg == nil {
		t.Fatal("Wrong data")
	}

	if seg.Left != left {
		t.Error("Wrong left")
	}

	if seg.Right != right {
		t.Error(" Wrong right")
	}

	if seg.Sum != sum {
		t.Error("Wrong sum")
	}
}
