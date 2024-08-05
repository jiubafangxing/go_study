	
package main
import(
	"log"
)
func test(){
	var t1, t2 interface{}
	t1,t2 = 1,2
	t1,t2 = []int{},[]int{} 
	log.Println(t1 == t2)
}

func main(){
	test()
}
