	
package main
import(
	"log"
)
func test(){
	type user struct{
		name string
		age int
	}
	users :=[...]user{
		{"laoli",12},
		{"laozhou",12},
	}
	log.Printf("%#v",users)
}

func main(){
	test()
}
