package chapter4
import(
	"fmt"
)
func CheckParam(a *int){
	fmt.Printf("the pointer is %p, target address %v, the value is %d\n", &a, a,*a);
}


func PassByValuemain(){
	o :=1
	a := &o;
	CheckParam(&o);
	fmt.Printf("the pointer is %p, target address %v, the value is %d\n", &a, a,*a);
}
