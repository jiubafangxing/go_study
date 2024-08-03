	
package main
import(
	"log"
	"strings"
)
func test(){
	s :=make([]string,100)
	for i:=0;i<100;i++{
		s[i] = "a"
	}
	result :=strings.Join(s,"")
	log.Println(result)
	
}

func main(){
	test()
}
