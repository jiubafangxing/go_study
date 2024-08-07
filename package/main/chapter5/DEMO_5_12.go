	
package chapter5
import(
	"log"
)
func DEMO_5_12test(){
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

func DEMO_5_12main(){
	DEMO_5_12test()
}
