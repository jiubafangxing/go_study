	
package main
import(
	"log"
)
func test(){
	a := [3]int{1,2,3}
	b := a
	b[1] = 4
	log.Println(a[1])
}

func main(){
	test()
}
