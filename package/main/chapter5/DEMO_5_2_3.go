	
package chapter5
import(
	"log"
)
func DEMO_5_2_3test(){
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

func DEMO_5_2_3main(){
	DEMO_5_2_3test()
}
