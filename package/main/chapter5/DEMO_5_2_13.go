	
package main
import(
	"log"
)
func test(){
	a := []int{1,2,3}
	a[1] = 3
	log.Println(a[1])
}

func main(){
	test()
}
