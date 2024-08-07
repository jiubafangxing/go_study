	
package chapter4
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
func DEMO_4_6_2test(a, b int )(int,error){
	if b == 0{
		return 0, DivErr{a,b}
	}
	return a/b, nil
}

func DEMO_4_6_2main(){
	result, err := DEMO_4_6_2test(1,0)
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
