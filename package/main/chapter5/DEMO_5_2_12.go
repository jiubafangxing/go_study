	
package main
import(
	"log"
	"unsafe"
	"reflect"
)
func test(){
	var a []int
	v := []int{}
	log.Println(a == nil, v == nil)
	log.Printf("a: %#v\n",(*reflect.SliceHeader)(unsafe.Pointer(&a)))
	log.Printf("v: %#v\n",(*reflect.SliceHeader)(unsafe.Pointer(&v)))
}

func main(){
	test()
}
