	
package main
import(
	"log"
	"bytes"
)
func test(){
	var b bytes.Buffer
	b.Grow(100)
	for i:=0;i<100; i++{
		b.WriteString(":")
	}
	result :=b.String()
	log.Println(result)
}

func main(){
	test()
}
