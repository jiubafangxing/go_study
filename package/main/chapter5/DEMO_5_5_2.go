	
package main
import(
	"log"
)
func test(){
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

func main(){
	test()
}
