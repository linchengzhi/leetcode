package recursion

func taijie(n int) int {
	if n <= 2 {
		return n
	}
	return taijie(n -1) + taijie(n -2)
}
