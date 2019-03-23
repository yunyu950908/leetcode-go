package golang

func findDiagonalOrder(matrix [][]int) []int {
	if !(len(matrix) != 0 && len(matrix[0]) != 0) {
		return []int{}
	}
	ir, ic := 0, 0
	lr, lc := len(matrix), len(matrix[0])
	result := make([]int, lr*lc)
	for i := range result {
		result[i] = matrix[ir][ic]
		if (ir+ic)%2 == 0 { // moving up
			if ic == lc-1 {
				ir += 1
			} else if ir == 0 {
				ic += 1
			} else {
				ic += 1
				ir -= 1
			}
		} else { // moving down
			if ir == lr-1 {
				ic += 1
			} else if ic == 0 {
				ir += 1
			} else {
				ir += 1
				ic -= 1
			}
		}
	}
	return result
}
