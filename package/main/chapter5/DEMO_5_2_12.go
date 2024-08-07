	
package chapter5
import(
	"log"
	"unsafe"
	"reflect"
)
func DEMO_5_2_12test(){
	var a []int
	v := []int{}
	log.Println(a == nil, v == nil)
	log.Printf("a: %#v\n",(*reflect.SliceHeader)(unsafe.Pointer(&a)))
	log.Printf("v: %#v\n",(*reflect.SliceHeader)(unsafe.Pointer(&v)))
}

func DEMO_5_2_12main(){
	DEMO_5_2_12test()
}
