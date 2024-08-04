	
package main
import(
	"log"
)
func test(){
	a :=[3]int{1,2,3}
	b :=a[0:2]
	log.Printf("b:%p",&b)
	b =b[0:1]
	log.Printf("b:%p",&b)
	b = append(b,4)
	log.Printf("b:%p",&b)
	b = append(b,5)
	log.Printf("b:%p",&b)
	b = append(b,6)
	log.Printf("b:%p",&b)
	b = append(b,7)
	log.Printf("b:%p",&b)

}

func main(){
	test()
}
