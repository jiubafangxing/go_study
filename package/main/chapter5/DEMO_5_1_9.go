	
package chapter5
import(
	"log"
	"strings"
)
func DEMO_5_1_9test(){
	s :=make([]string,100)
	for i:=0;i<100;i++{
		s[i] = "a"
	}
	result :=strings.Join(s,"")
	log.Println(result)
	
}

func DEMO_5_1_9main(){
	DEMO_5_1_9test()
}
