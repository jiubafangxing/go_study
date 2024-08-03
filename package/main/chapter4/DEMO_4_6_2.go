	
package main
import(
	"fmt"
	"log"
)
type DivErr struct {
    x , y int

}
func (DivErr) Error()(string){
	return "division by zero"
}
func test(a, b int )(int,error){
	if b == 0{
		return 0, DivErr{a,b}
	}
	return a/b, nil
}

func main(){
	result, err := test(1,0)
	if nil != err{
		switch e:= err.(type){
			case  DivErr:
				fmt.Println(e, e.x,e.y)
			default:
				fmt.Println(e)
		}
		log.Fatalln(err)
	}
	fmt.Println(result)

}
