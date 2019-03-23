package golang

func plusOne(digits []int) []int {
	carry := 1
	idx := len(digits) - 1
	for ; carry == 1 && idx >= 0; idx -= 1 {
		if carry == 1 {
			digits[idx] += 1
			carry = 0
		}
		if digits[idx] > 9 {
			carry = 1
		}
		digits[idx] = digits[idx] % 10
	}
	if carry == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
