	
package chapter5
import(
	"log"
)
func DEMO_5_1_1test(){
	s := "你好\x61"
	log.Printf("%s\n",s)
}

func DEMO_5_1_1main(){
	DEMO_5_1_1test()
}
