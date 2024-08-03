	
package main
import(
	"log"
)
func test(){
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

func main(){
	test()
}
