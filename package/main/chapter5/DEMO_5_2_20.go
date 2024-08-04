	
package main
import(
	"log"
)
func test(){
	b := make([]byte,3,5)
	n := copy(b,"abcd")
	log.Println(n,b)
}

func main(){
	test()
}
