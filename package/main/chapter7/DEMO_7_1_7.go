	
package chapter7
import(
	"log"
)
type item interface{
	name1()(string)
}

type user struct{
	name string
}

func (u user) name1()string{
	return u.name
}
func DEMO_7_1_7test(){
	var u user = user{
		name:"laoli",
	}
	var i item = u
	log.Println(i.(item).name1())
}

func DEMO_7_1_7main(){
	DEMO_7_1_7test()
}
