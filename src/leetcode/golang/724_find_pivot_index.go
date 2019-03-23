package golang

func pivotIndex(nums []int) int {
	sum, leftSum := 0, 0
	for _, v := range nums {
		sum += v
	}
	for i, v := range nums {
		if leftSum == sum-leftSum-v {
			return i
		}
		leftSum += v
	}
	return -1
}
