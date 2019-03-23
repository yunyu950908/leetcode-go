package golang

func spiralOrder(matrix [][]int) []int {
	if !(len(matrix) != 0 && len(matrix[0]) != 0) {
		return []int{}
	}
	ir, ic := 0, 0
	lr, lc := len(matrix), len(matrix[0])
	result := make([]int, lr*lc)
	// 1 (left -> right),
	// 2 (up -> down),
	// 3 (right -> left),
	// 4 (down -> up)
	direction := 1
	boundary := ic
	for i := range result {
		switch direction {
		case 1:
			result[i] = matrix[ir][boundary]
			boundary += 1
			if boundary == lc {
				direction = 2
				ir += 1
				boundary = ir
			}
		case 2:
			result[i] = matrix[boundary][lc-1]
			boundary += 1
			if boundary == lr {
				direction = 3
				lc -= 1
				boundary = lc - 1
			}
		case 3:
			result[i] = matrix[lr-1][boundary]
			boundary -= 1
			if boundary == ic-1 {
				direction = 4
				lr -= 1
				boundary = lr - 1
			}
		case 4:
			result[i] = matrix[boundary][ic]
			boundary -= 1
			if boundary == ir-1 {
				direction = 1
				ic += 1
				boundary = ic
			}
		}
	}
	return result
}
