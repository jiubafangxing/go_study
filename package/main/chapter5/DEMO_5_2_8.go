	
package main
import(
	"log"
)
func test(a *[3]int){
	a[1] = 1
}

func main(){
	a := [...]int{2,2,2}
	test(&a)
	log.Println(a[1])
}
