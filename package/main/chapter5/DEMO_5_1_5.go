	
package chapter5
import(
	"log"
	"reflect"
	"unsafe"
)
func DEMO_5_1_5test(){
	s := "abcdef"
	s1 :=s[:3]
	log.Println(s1)
	log.Printf("head address %#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	log.Printf("head address %#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
}

func DEMO_5_1_5main(){
	DEMO_5_1_5test()
}
