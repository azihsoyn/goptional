package goptional_test

import (
	"testing"

	"github.com/azihsoyn/goptional"
)

func BenchmarkGetOrElse(b *testing.B) {
	var n *int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goptional.GetOrElse(n, 0)
	}
}

func BenchmarkString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goptional.String("string")
	}
}
