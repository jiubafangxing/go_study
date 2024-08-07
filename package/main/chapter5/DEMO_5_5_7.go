	
package chapter5
import(
	"log"
	"unsafe"
)
func DEMO_5_5_7test(){
	var a struct{}
	var b [100]struct{}
	log.Println(unsafe.Sizeof(a),unsafe.Sizeof(b))
}

func DEMO_5_5_7main(){
	DEMO_5_5_7test()
}
