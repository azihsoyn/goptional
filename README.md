# goptional
goptional is treat pointer type as optional (using minimum reflect)

# Installation
`go get github.com/azihsoyn/goptional` or `go get gopkg.in/azihsoyn/goptional.v1`

# Usage
```go
package main

import (
	"fmt"

	"github.com/azihsoyn/goptional"
)

func main() {
	var np *int
	// fmt.Println(*n) this cause panic
	fmt.Println(goptional.GetOrElse(np, 0))
	// Output: 0

	n := 1
	np = &n
	fmt.Println(goptional.GetOrElse(np, 0))
	// 1
}
```

# benchmark
```
$ go test -bench  . -benchmem
BenchmarkGetOrElse-8       50000000            22.4 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/azihsoyn/goptional    1.203s
```
