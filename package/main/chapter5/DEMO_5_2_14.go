	
package chapter5
import(
	"log"
)
func DEMO_5_2_14test(){
	x :=[][]int{
		{1,2,3},
		{10},
		{100,200},
	}
	log.Println(x[1])
	x[2] = append(x[2],200)
	x[2] = append(x[2],300)
	x[2] = append(x[2],400)
	log.Println(x[2])
}

func DEMO_5_2_14main(){
	DEMO_5_2_14test()
}
