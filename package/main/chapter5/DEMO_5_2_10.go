	
package main
import(
	"log"
)
func test(){
	s1 := make([]int,3,5)
	s2 := make([]int , 3)
	s3 := []int{10,29,5:1}
	log.Println(s1,len(s1), cap(s1))
	log.Println(s2,len(s2), cap(s2))
	log.Println(s3,len(s3), cap(s3))
}

func main(){
	test()
}
