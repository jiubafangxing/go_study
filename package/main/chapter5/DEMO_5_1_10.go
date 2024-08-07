	
package chapter5
import(
	"log"
	"bytes"
)
func DEMO_5_1_10test(){
	var b bytes.Buffer
	b.Grow(100)
	for i:=0;i<100; i++{
		b.WriteString(":")
	}
	result :=b.String()
	log.Println(result)
}

func DEMO_5_1_10main(){
	DEMO_5_1_10test()
}
