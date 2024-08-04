	
package main
import(
	"sync"
)
type data struct{
	sync.Mutex
	buf [1024]byte
}
func test(){
	d := data{}
	d.Lock()
	defer d.Unlock()
}

func main(){
	test()
}
