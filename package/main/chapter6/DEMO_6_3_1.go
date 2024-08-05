	
package main
import(
	"log"
	"fmt"
	"reflect"
)
type S struct{}

type T struct{
	S
}

func (*S) printS1(){
	log.Println("s1")
}
func (S) printS(){
	log.Println("s")
}

func (*T) printT1(){
	log.Println("t1")
}

func (T) printT(){
	log.Println("t")
}

func test(a interface{}){
	t :=reflect.TypeOf(a)
	log.Println("methods,",t.NumMethod())
	for i,n :=0, t.NumMethod();i<n;i++{
		m :=t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main(){
	var t T
	test(t)
	test(&t)
}
