	
package chapter8
import(
	"log"
	"sync"
	"math"
	"runtime"
)
func count(){
	x :=0
	for i:=0;i< math.MaxUint32;i++{
		x +=i
	}
	log.Println(x)
}
func DEMO_8_1_6test2(n int){
	for i:=0; i<n; i++{
		count()
	}	
}

func DEMO_8_1_6test( n int ){
	var wg sync.WaitGroup
	for i:=0;i< n;i++{
		x := i
		wg.Add(1)
		go func(){
			log.Printf("线程%d",x)
			count()
			wg.Done()
		}()
	}
	wg.Wait()
}

func DEMO_8_1_6main(){
	n := runtime.GOMAXPROCS(0)
	DEMO_8_1_6test2(n)
}
