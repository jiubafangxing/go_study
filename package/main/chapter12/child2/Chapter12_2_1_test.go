package child2

import (
	"log"
	"testing"
)

func C1221Add(x, y int) int {
	return x + y
}

func BenchmarkC1221Add(b *testing.B) {
	log.Printf("benchmark b.N:%d\n", b.N)
	for i := 0; i < b.N; i++ {
		_ = C1221Add(1, 2)
	}
}
