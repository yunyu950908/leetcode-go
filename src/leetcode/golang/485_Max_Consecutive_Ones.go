package golang

func findMaxConsecutiveOnes(nums []int) int {
	var max, count int
	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
	}
	return max
}
