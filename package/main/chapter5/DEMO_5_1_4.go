	
package main
import(
	"log"
)
func test(){
	s := "ab"+
		"cd"
	log.Println(s == "abcd")
}

func main(){
	test()
}
