package golang

func dominantIndex(nums []int) int {
	max, second, idx := 0, 0, 0
	for i, v := range nums {
		if v > max {
			idx, max = i, v
		}
	}
	for i, v := range nums {
		if v > second && i != idx {
			second = v
		}
	}
	if max >= second*2 {
		return idx
	}
	return -1
}
