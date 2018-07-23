package interval

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	a := &Array{}
	a = a.Set(0, 150, 0)
	a = a.Set(10, 20, 10)
	fmt.Println(a)
	a = a.Set(20, 30, 10)
	fmt.Println(a)
	a = a.Max(5, 40, 5)
	fmt.Println(a)
}
