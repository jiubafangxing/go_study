package chapter5
import(
	"log"
)
func DEMO_5_4_8test(i map[string]int){
	log.Printf("x: %p\n",i)
}
func DEMO_5_4_8main(){
	dict1 := map[string]int{}
	DEMO_5_4_8test(dict1)
	log.Printf("x: %p\n",dict1)

}
