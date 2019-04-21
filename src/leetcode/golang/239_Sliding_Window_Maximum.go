package golang

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	window, res := make([]int, 0), make([]int, 0)
	for i, v := range nums {
		if i >= k && i-k >= window[0] {
			window = window[1:]
		}
		for len(window) > 0 && v >= nums[window[len(window)-1]] {
			window = window[0 : len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}
