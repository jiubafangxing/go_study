	
package chapter4
import(
	"log"
)
func DEMO_4_6_5test(){
	defer func(){
		if err := recover(); nil != err{
			log.Fatalln(err)
		}	
	}()	
	defer func(){
		panic("you are dead")
	}()
	panic("i am dead")
	
}

func DEMO_4_6_5main(){
	DEMO_4_6_5test()
}
