	
package main
import(
	"log"
)
func test(){
	var a byte
	a = 'a'
	ra := rune(a)
	sa := string(a)
	log.Println(ra)
	log.Println(sa)
}

func main(){
	test()
}
