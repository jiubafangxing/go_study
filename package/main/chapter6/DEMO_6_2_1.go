	
package chapter6
import(
	"sync"
)
type data struct{
	sync.Mutex
	buf [1024]byte
}
func DEMO_6_2_1test(){
	d := data{}
	d.Lock()
	defer d.Unlock()
}

func DEMO_6_2_1main(){
	DEMO_6_2_1test()
}
