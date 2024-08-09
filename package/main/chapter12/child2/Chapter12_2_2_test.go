package child2

import (
	"testing"
	"time"
)

func sleep() {
	time.Sleep(time.Second)
}

func BenchmarkSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sleep()
	}
}
