	
package main
import(
	"log"
	"unsafe"
)
func test(){
	var a struct{}
	var b [100]struct{}
	log.Println(unsafe.Sizeof(a),unsafe.Sizeof(b))
}

func main(){
	test()
}
