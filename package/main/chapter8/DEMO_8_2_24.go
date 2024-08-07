	
package chapter8
import(
	"sync"
)
// this will be dead lock 
func DEMO_8_2_24test(){
	var m sync.Mutex
	m.Lock()
	{
		m.Lock()
		m.Unlock()
	}
	m.Unlock()
}

func DEMO_8_2_24main(){
	DEMO_8_2_24test()
}
