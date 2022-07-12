package _00_199

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	return analyzeMaxPoints(0, points, 0) + 2
}

func analyzeMaxPoints(start int, points [][]int, max int) int {
	if len(points)-start <= max {
		return max
	}
	for i := start + 1; i < len(points); i++ {
		d := countLineNum(start, i, points)
		if d > max {
			max = d
		}
		if len(points)-i <= max {
			break
		}
	}
	return analyzeMaxPoints(start+1, points, max)
}

func countLineNum(start, next int, arr [][]int) int {
	var num = 0
	dx := arr[next][0] - arr[start][0]
	dy := arr[next][1] - arr[start][1]
	for i := next + 1; i < len(arr); i++ {
		dx2 := arr[i][0] - arr[start][0]
		dy2 := arr[i][1] - arr[start][1]
		if dx*dy2 != dy*dx2 {
			continue
		}
		num++
	}
	return num
}
