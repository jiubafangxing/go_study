	
package chapter5
import(
	"log"
)
func DEMO_5_4_6test(){
	dict := make(map[int]int, 10)
	for i:=0;i<10;i++{
		dict[i] =i
	}
	for k,_ := range dict{
		delete(dict,k)

		log.Println(dict)
	}
}

func DEMO_5_4_6main(){
	DEMO_5_4_6test()
}
