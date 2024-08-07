	
package chapter4
import(
	"log"
)
func DEMO_4_6_4test(){
 	defer func(){
		log.Println("DEMO_4_6_4test1")
	}()
 	defer func(){
		log.Println("DEMO_4_6_4test2")
	}()
	panic("outter receive")
}

func DEMO_4_6_4main(){
	defer func(){
		if err :=recover(); nil != err{
			log.Fatalln(err)
		}
	}()
	DEMO_4_6_4test()

}
