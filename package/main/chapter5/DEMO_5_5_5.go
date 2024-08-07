	
package chapter5
import(
	"log"
)
func DEMO_5_5_5test(){
	type data struct{
		x int
		y map[string]int
	}
	d1 := data{
		x:1,
	}
	d2 := data{
		x:2,
	}
	log.Println(d1 == d2)

}

func DEMO_5_5_5main(){
	DEMO_5_5_5test()
}
