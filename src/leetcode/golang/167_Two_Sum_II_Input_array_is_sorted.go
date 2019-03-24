package golang

func twoSum(numbers []int, target int) []int {
	var in1, in2, sum int
	in1, in2 = 0, len(numbers)-1
	for in1 < in2 {
		sum = numbers[in1] + numbers[in2]
		if sum > target {
			in2--
		} else if sum < target {
			in1++
		} else {
			return []int{in1 + 1, in2 + 1}
		}
	}
	return []int{}
}
