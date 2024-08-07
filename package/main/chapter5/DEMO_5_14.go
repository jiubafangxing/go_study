	
package chapter5
import(
	"log"
	"reflect"
)
func DEMO_5_14test(){
	type user struct{
		name string `姓名`
		age int	`年龄`
	}
	user1  := user{
		"laoli",
		12,
	}
	v :=reflect.ValueOf(user1)
	t :=v.Type()
	for i,n :=0,t.NumField();i<n;i++{
		log.Printf("%s: %v\n",t.Field(i).Tag,v.Field(i))
	}
}

func DEMO_5_14main(){
	DEMO_5_14test()
}
