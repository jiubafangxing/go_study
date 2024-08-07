	
package chapter5
import(
	"log"
)
func DEMO_5_5_6test(){
	type user struct{
		name string
		age int
	}
	u := &user{
		name:"wanggelun",
		age:12,
	}
	log.Println(u)
	up := &u
	(*up).name = "kkek"
}


func DEMO_5_5_6main(){
	DEMO_5_5_6test()
}
