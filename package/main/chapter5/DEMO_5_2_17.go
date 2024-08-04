	
package main
import(
	"log"
)
func test(){
	a :=[5]int{1,2,3,4,5} 
	b :=a[0:2:5]
	c := b[0:1]
	log.Println(cap(c))
}

func main(){
	test()
}
