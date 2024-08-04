	
package main
import(
	"log"
)
func test(){
	var dict1 map[int]int
	m2 := make(map[int]int,1)
	log.Println(dict1 == nil)
	log.Println(m2 == nil)
}

func main(){
	test()
}
