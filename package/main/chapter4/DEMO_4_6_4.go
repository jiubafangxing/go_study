	
package main
import(
	"log"
)
func test(){
 	defer func(){
		log.Println("test1")
	}()
 	defer func(){
		log.Println("test2")
	}()
	panic("outter receive")
}

func main(){
	defer func(){
		if err :=recover(); nil != err{
			log.Fatalln(err)
		}
	}()
	test()

}
