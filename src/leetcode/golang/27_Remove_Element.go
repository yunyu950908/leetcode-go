package golang

func removeElement(nums []int, val int) int {
	ln := len(nums)
	for i := 0; i < ln; i++ {
		if nums[i] == val {
			nums[i] = nums[ln-1]
			ln--
			i--
		}
	}
	return ln
}
