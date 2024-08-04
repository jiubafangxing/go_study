	
package main
import(
	"log"
)
func test(){
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


func main(){
	test()
}
