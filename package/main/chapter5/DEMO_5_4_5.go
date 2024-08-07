	
package chapter5
import(
	"log"
)
func DEMO_5_4_5test(){
	var dict1 map[int]int
	m2 := make(map[int]int,1)
	log.Println(dict1 == nil)
	log.Println(m2 == nil)
}

func DEMO_5_4_5main(){
	DEMO_5_4_5test()
}
