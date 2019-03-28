package golang

import (
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	var sum, l, r, ln int
	ln = len(nums)
	minLen := ln + 1
	for l < ln {
		if r < ln && sum < s {
			sum += nums[r]
			r++
		} else {
			sum -= nums[l]
			l++
		}

		if sum >= s {
			minLen = int(math.Min(float64(minLen), float64(r-l)))
		}
	}
	if minLen == ln+1 {
		return 0
	}
	return minLen
}

//func minSubArrayLen(s int, nums []int) int {
//	var i, j, k, sum, ans int
//	ln := len(nums)
//	max := int(^uint(0) >> 1)
//	ans = max
//	for i = 0; i < ln; i++ {
//		for j = i; j < ln; j++ {
//			sum = 0
//			for k = i; k <= j; k++ {
//				sum += nums[k]
//			}
//			if sum >= s {
//				ans = int(math.Min(float64(ans), float64(j-i+1)))
//				break
//			}
//		}
//	}
//	if ans == max {
//		return 0
//	} else {
//		return ans
//	}
//}
