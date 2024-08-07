	
package chapter5
import(
	"log"
)

type user struct{
	name string
	age int
}
func DEMO_5_4_4test(){
	dict :=map[int]user{
		1:{"lao1l",1},
		2:{"manba",2},
	}
	user1 :=dict[1]
	user1.name = "naruto"
	dict[1]= user1
	log.Println(dict)
}

func DEMO_5_4_4main(){
	DEMO_5_4_4test()
}
