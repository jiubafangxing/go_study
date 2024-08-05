	
package main
import(
	"log"
	"sync"
)
type data [8]struct{
	id int
	result int
}

func test(){
	var wg sync.WaitGroup
	var data1 data
	//simulate local storage
	for i:=0;i<8;i++{
		wg.Add(1)
	 	go func(x int){
			data1[x].id = x
			data1[x].result = x	
			wg.Done()
		}(i)	
	}
	wg.Wait()
	log.Println("%+v\n",data1)	
}

func main(){
	test()
}
