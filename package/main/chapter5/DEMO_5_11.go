	
package chapter5
import(
	"log"
)
func DEMO_5_11test(){
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

func DEMO_5_11main(){
	DEMO_5_11test()
}
