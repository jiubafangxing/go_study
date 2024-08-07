	
package chapter5
import(
	"log"
)
func DEMO_5_13test(){
	type file struct{
		name string
	}
	type log2 struct{
		name string
	}

	type data struct{
		file
		log2
	}
	data1 :=data{
		
	}
	data1.log2.name = "1"
	data1.file.name = "xx"
	log.Println(data1)

}

func DEMO_5_13main(){
	DEMO_5_13test()
}
