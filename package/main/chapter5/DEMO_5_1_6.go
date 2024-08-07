	
package chapter5
import(
	"log"
)
func DEMO_5_1_6test(){
	var a string 
	a = "世界你好"
	for i :=0; i< len(a);i++{
		log.Printf("[%c]",a[i])
	}
	for _,c := range a{


		log.Printf("[%c]",c)
	}
}

func DEMO_5_1_6main(){
	DEMO_5_1_6test()
}
