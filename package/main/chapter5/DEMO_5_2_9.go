	
package main
import(
	"log"
)
func test(){
	m := map[string][2]int{
		"a":{1,2},
	}
	subarr := m["a"]
	log.Printf("%#v",subarr)
	item :=m["a"][:]
	log.Printf("%#v",item)
}

func main(){
	test()
}
