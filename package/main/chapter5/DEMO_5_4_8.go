package main
import(
	"log"
)
func test(i map[string]int){
	log.Printf("x: %p\n",i)
}
func main(){
	dict1 := map[string]int{}
	test(dict1)
	log.Printf("x: %p\n",dict1)

}
