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

func Example_string() {
	fmt.Println(*goptional.String("hello"))
	// Output: hello
}

func Example_int() {
	fmt.Println(*goptional.Int(1))
	// Output: 1
}

func Example_int64() {
	fmt.Println(*goptional.Int64(1))
	// Output: 1
}

func Example_bool() {
	fmt.Println(*goptional.Bool(true))
	// Output: true
}
