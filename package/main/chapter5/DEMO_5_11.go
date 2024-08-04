	
package main
import(
	"log"
)
func test(){
	type data struct{
	 *int
	 string
	}
	x:=10
	data1 := data{
		&x,
		"asdf",
	}
	log.Println(data1)
	log.Println(*data1.int)
}

func main(){
	test()
}
