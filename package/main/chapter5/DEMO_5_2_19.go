	
package chapter5
import(
	"log"
)
func DEMO_5_2_19test(){
	s :=[]int{12,2,34,3,23,112}
	s2 := s[3:5]
	s3 := copy(s,s2)
	log.Println(s3, s2)
	log.Println(s)
}

func DEMO_5_2_19main(){
	DEMO_5_2_19test()
}
