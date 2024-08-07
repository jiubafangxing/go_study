	
package chapter8
import(
	"log"
	"sync"
	"time"
)

type data struct{
	sync.Mutex
}
// this is error , can not lock
//func (d data) p(s string){
func (d *data) p(s string){
	d.Lock()
	defer d.Unlock()
	for i:=0;i< 5;i++{
		log.Println(s)
		time.Sleep(time.Second)	
	}
}
func DEMO_8_2_23test(){
	var wg sync.WaitGroup
	wg.Add(2)
	var d1 data = data{}
	go func(){
		defer wg.Done()
		d1.p("a")
	}()
	go func(){
		defer wg.Done()
		d1.p("b")
	}()
	wg.Wait()
}

func DEMO_8_2_23main(){
	DEMO_8_2_23test()
}
