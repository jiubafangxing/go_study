	
package main
import(
	"log"
)
func test(){
	a := [...][4]int{
		{1,2,3,4},
		{5,56,6,7},
	}
	log.Printf("#%v",a)
}

func main(){
	test()
}
