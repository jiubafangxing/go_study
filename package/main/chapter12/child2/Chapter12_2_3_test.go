package child2

import(
	"testing"
	"time"
)
func C1223Add(a , b int)(int){
	return a+b
}

func BenchmarkC1223Add(b *testing.B){
	time.Sleep(time.Second)
	b.ResetTimer()
	for i:=0; i< b.N;i++{
		_ = C1223Add(1,2)
	}
}

