	
package chapter5
import(
	"log"
)
func DEMO_5_1_12test(){
	var a byte
	a = 'a'
	ra := rune(a)
	sa := string(a)
	log.Println(ra)
	log.Println(sa)
}

func DEMO_5_1_12main(){
	DEMO_5_1_12test()
}
