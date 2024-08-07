	
package chapter5
import(
	"log"
)
func DEMO_5_4_3test(){
	dict := map[string]int{}
	for i:=0;i<10;i++{
		dict[string('a'+i)] = i
	}
	for s :=0;s<5; s++{
		for k,v := range dict{
			log.Printf("key: %s, value: %d,", k,v)
		}
		log.Println()
	}
}

func DEMO_5_4_3main(){
	DEMO_5_4_3test()
}
