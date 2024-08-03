	
package main
import "log"
func test()(z int){
	defer func(){
		z +=100
	}()
	return 100
}

func main(){
	result := test()
	log.Println("result is ", result)
}
