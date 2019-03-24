package golang

import "sort"

func arrayPairSum(nums []int) int {
	var sum int
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
