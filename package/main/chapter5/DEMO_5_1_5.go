	
package main
import(
	"log"
	"reflect"
	"unsafe"
)
func test(){
	s := "abcdef"
	s1 :=s[:3]
	log.Println(s1)
	log.Printf("head address %#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	log.Printf("head address %#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
}

func main(){
	test()
}
