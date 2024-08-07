	
package chapter5
import(
	"log"
)
func DEMO_5_2_6test(){
	x, y := 2,3
	a :=[]*int{&x, &y}
	b := &a
	log.Printf("type:%T, v:%v\n",a,a)
	log.Printf("type:%T, v:%v\n",b,b)
	log.Println(*((*b)[0]))
}

func DEMO_5_2_6main(){
	DEMO_5_2_6test()
}
