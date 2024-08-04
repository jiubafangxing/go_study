	
package main
import(
	"log"
)
func test(){
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

func main(){
	test()
}
