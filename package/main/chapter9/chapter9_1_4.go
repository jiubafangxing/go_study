package chapter9
import(
	"fmt"
	"log"
)
type C914_A struct {
	x int
}
type C914_B struct {
	x string
}
func (c C914_B)c914print()(string){
	return fmt.Printf("%s",c.x)	
}
func (c C914_A)c914print()(string){
	return fmt.Printf("%d",c.x)	
}
type C914_AB interface{
	C914_A | C914_B
	c914print()
}
func print914[T C914_AB](t T){
	log.Println(t.c914print())
}
func Chapter914Test() {
	c1 := C914_A{
		x:1,
	}
	c2 := C914_B{
		x:"hello ",
	}
	print914(c1)
	print914(c2)
	print914("a")
	return
}
