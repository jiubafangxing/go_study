	
package main
import(
	"log"
)
func test(){
	s := "abcd"	
	a := []byte(s)
	a[0] = 'c'
	news :=string(a)
	log.Printf(news)
}

func main(){
	test()
}
