package chapter9
import(
	"log"
)
type C912Data[T any] struct{
	x T
}

func (c C912Data[T]) print(){
	log.Println(c.x)	

}
func Chapter9_1_2Test() {
	c1 := C912Data[int]{
		x: 912,
	}
	c1.print()
	c2 := C912Data[string]{
		x: "hello generics",
	}
	c2.print()
	return
}
