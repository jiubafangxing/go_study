	
package chapter5
import(
	"log"
)
func DEMO_5_2_9test(){
	m := map[string][2]int{
		"a":{1,2},
	}
	subarr := m["a"]
	log.Printf("%#v",subarr)
	item :=m["a"][:]
	log.Printf("%#v",item)
}

func DEMO_5_2_9main(){
	DEMO_5_2_9test()
}
