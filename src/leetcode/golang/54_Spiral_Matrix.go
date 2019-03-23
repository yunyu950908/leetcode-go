package golang

func spiralOrder(matrix [][]int) []int {
	if !(len(matrix) != 0 && len(matrix[0]) != 0) {
		return []int{}
	}
	var ir, ic, nr, nc int
	lr, lc := len(matrix), len(matrix[0])
	result := make([]int, lr*lc)
	// init seen
	seen := make([][]bool, lr)
	for i := range seen {
		seen[i] = make([]bool, lc)
	}
	// direction
	dr := [4]int{0, 1, 0, -1}
	dc := [4]int{1, 0, -1, 0}
	di := 0
	for i := range result {
		result[i] = matrix[ir][ic]
		seen[ir][ic] = true
		nr, nc = ir+dr[di], ic+dc[di]
		if nr >= 0 && nr < lr && nc >= 0 && nc < lc && !seen[nr][nc] {
			ir, ic = nr, nc
		} else {
			di = (di + 1) % 4
			ir, ic = ir+dr[di], ic+dc[di]
		}
	}
	return result
}
