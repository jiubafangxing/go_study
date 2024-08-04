	
package main
import(
	"log"
)
func test(){
	var b [100]struct{}
	log.Println(cap(b),len(b))
}

func main(){
	test()
}
