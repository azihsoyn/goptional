package goptional_test

import (
	"fmt"

	"github.com/azihsoyn/goptional"
)

func Example_getOrElse() {
	var np *int
	// fmt.Println(*n) this cause panic
	fmt.Println(goptional.GetOrElse(np, 0))
	// Output: 0

	n := 1
	np = &n
	fmt.Println(goptional.GetOrElse(np, 0))
	// 1
}
