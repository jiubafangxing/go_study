	
package main
import(
	"log"
)
func test(){
	var a interface{} = nil
	var b interface{}= (*int)(nil)
	log.Println(a == nil)
	log.Println(b == nil)
}

func main(){
	test()
}
