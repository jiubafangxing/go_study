	
package main
import(
	"log"
)
func test(){
	defer func(){
		if err := recover() ; nil != err{
			log.Fatalln(err)
		}
	}()

	panic("i am dead")
	log.Println("still run ")
}

func main(){
	test()
}
