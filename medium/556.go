package medium

import (
	"math"
	"sort"
)

func nextGreaterElement(n int) int {
	if n < 10 {
		return -1
	}
	inputs := []int{}
	for n > 0 {
		r := n % 10
		inputs = append(inputs, r)
		n = n / 10
	}

	tp := 1
	for tp < len(inputs) && inputs[tp] >= inputs[tp-1] {
		tp++
	}
	if tp == len(inputs) {
		return -1
	}

	// index: 5 4 3 2 1 0
	// n:     2 3 0 2 4 1
	//              ^
	//             tp

	subArray := inputs[:tp]
	sort.Ints(subArray)
	for i, x := range subArray {
		if x > inputs[tp] {
			subArray[i] = inputs[tp]
			inputs[tp] = x
			break
		}
	}
	answer := 0
	for i := len(inputs) - 1; i >= tp; i-- {
		answer = answer*10 + inputs[i]
	}
	for i := 0; i < len(subArray); i++ {
		answer = answer*10 + subArray[i]
	}
	if answer > math.MaxInt32 {
		return -1
	}
	return answer
}
