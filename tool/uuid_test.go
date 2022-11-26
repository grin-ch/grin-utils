package tool

import (
	"testing"
)

func BenchmarkMustUUIDv4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MustUUIDv4()
	}
}
