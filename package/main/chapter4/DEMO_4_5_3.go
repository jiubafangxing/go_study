	
package chapter4
import "log"
func DEMO_4_5_3test()(z int){
	defer func(){
		z +=100
	}()
	return 100
}

func DEMO_4_5_3main(){
	result := DEMO_4_5_3test()
	log.Println("result is ", result)
}
