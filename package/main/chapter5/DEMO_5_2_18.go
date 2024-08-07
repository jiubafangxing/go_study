	
package chapter5
import(
	"log"
)
func DEMO_5_2_18test(){
	a :=[3]int{1,2,3}
	b :=a[0:2]
	log.Printf("b:%p",&b)
	b =b[0:1]
	log.Printf("b:%p",&b)
	b = append(b,4)
	log.Printf("b:%p",&b)
	b = append(b,5)
	log.Printf("b:%p",&b)
	b = append(b,6)
	log.Printf("b:%p",&b)
	b = append(b,7)
	log.Printf("b:%p",&b)

}

func DEMO_5_2_18main(){
	DEMO_5_2_18test()
}
