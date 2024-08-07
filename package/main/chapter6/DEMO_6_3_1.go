	
package chapter6
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

func DEMO_6_3_1test(a interface{}){
	t :=reflect.TypeOf(a)
	log.Println("methods,",t.NumMethod())
	for i,n :=0, t.NumMethod();i<n;i++{
		m :=t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func DEMO_6_3_1main(){
	var t T
	DEMO_6_3_1test(t)
	DEMO_6_3_1test(&t)
}
