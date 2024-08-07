	
package chapter5
import(
	"log"
)
func DEMO_5_2_17test(){
	a :=[5]int{1,2,3,4,5} 
	b :=a[0:2:5]
	c := b[0:1]
	log.Println(cap(c))
}

func DEMO_5_2_17main(){
	DEMO_5_2_17test()
}
