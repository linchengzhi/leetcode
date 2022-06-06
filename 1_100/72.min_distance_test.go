package __100

import (
	"fmt"
	"testing"
)

func Test_minDistance(t *testing.T) {
	got := minDistance("abc", "cagd")
	fmt.Println(got)
	got = minDistance("a", "")
	fmt.Println(got)
}
