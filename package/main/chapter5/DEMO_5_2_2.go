	
package main
import(
	"log"
)
func test(){
	var a [4]int
	b := [3]int{}
	c := [4]int{2:1}
	d := [...]int{2:1}
	e := [...]int{1,2,3}
	log.Println(len(a))
	log.Println(len(b))
	log.Println(len(c))
	log.Println(len(d))
	log.Println(len(e))
}

func main(){
	test()
}
