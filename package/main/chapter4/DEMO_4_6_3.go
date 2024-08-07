	
package chapter4
import(
	"log"
)
func DEMO_4_6_3test(){
	defer func(){
		if err := recover() ; nil != err{
			log.Fatalln(err)
		}
	}()

	panic("i am dead")
	log.Println("still run ")
}

func DEMO_4_6_3main(){
	DEMO_4_6_3test()
}
