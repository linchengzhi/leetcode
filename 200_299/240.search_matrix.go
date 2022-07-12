package _00_299

func searchMatrix(matrix [][]int, target int) bool {
	var i, j = len(matrix) - 1, 0
	for ; i >= 0; i-- {
		for j <= len(matrix[i])-1 {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {
				break
			}
			j++
		}
	}
	return false
}
