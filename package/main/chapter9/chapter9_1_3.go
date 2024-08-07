package chapter9
import(
	"fmt"
	"log"
)
type C913Data struct{
	x int
}

func (c C913Data) String()string{
	return fmt.Sprintf("%d", c.x)	
}

func printGenerics[T fmt.Stringer](v T){
	log.Println(v)	
}
func Chapter913Test() {
	printGenerics(C913Data{x:3,})
	return
}
