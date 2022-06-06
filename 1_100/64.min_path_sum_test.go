package __100

import "testing"

func TestMinPathSum(t *testing.T) {
	grid := make([][]int, 0)
	grid = append(grid, []int{1,3,1})
	grid = append(grid, []int{1,5,1})
	grid = append(grid, []int{4,2,1})
	d := minPathSum(grid)
	t.Log(d)
}