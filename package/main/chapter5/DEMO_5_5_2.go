	
package chapter5
import(
	"log"
)
func DEMO_5_5_2test(){
	type user struct{
		name string
		age int
	}
	u1 := user{
		"wang",
		12,
	}
	u2 := user{
		"wang",
	}
	log.Println(u1,u2)
}

func DEMO_5_5_2main(){
	DEMO_5_5_2test()
}
