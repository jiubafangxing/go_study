	
package chapter5
import(
	"log"
)
func DEMO_5_4_1test(){
	dict :=  make(map[string]int)
	dict["a"] = 1
	dict["1"]= 2
	log.Println(dict)

	phoneNum := map[int]struct{
		x string	
	}{
		1:{"13635235578"},
	}
	log.Println(phoneNum[1])
}

func DEMO_5_4_1main(){
	DEMO_5_4_1test()
}
