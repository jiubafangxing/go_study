	
package main
import(
	"log"
)
func test(){
	type file struct{
		name string
	}
	type data struct{
		file
		name string
	}
	data1 := data{
		name:"lili",
	}
	data1.file.name = "huoli"
	log.Println(data1.name)
	log.Println(data1.file.name)

}

func main(){
	test()
}
