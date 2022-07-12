package __99

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := IntsToList([]int{2, 4, 3})
	l2 := IntsToList([]int{5, 6, 4})
	result := ListToInt(addTwoNumbers(l1, l2))
	fmt.Println(result)
}
