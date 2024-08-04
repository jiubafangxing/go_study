	
package main
import(
	"log"
)
func test(){
	s :=[]int{12,2,34,3,23,112}
	s2 := s[3:5]
	s3 := copy(s,s2)
	log.Println(s3, s2)
	log.Println(s)
}

func main(){
	test()
}
