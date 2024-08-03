	
package main
import(
	"log"
)
func test(){
	var a []byte
	a = append(a,"hello"...)
	a = append(a,'a','b')
	log.Println(a[0])
}

func main(){
	test()
}
