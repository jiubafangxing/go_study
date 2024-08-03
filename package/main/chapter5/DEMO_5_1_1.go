	
package main
import(
	"log"
)
func test(){
	s := "你好\x61"
	log.Printf("%s\n",s)
}

func main(){
	test()
}
