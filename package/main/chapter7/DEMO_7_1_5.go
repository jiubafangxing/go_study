	
package main
import(
	"log"
)

type data1 struct{
	
}

func (data1) toString()(string){
	return "data1"
}

type node struct{
	data interface{
		toString()(string)
	}
}
func test(){
	var d interface{
		toString()(string)
	} = data1{}
	var node1 node = node{
		data:d,
	}
	log.Println(node1.data.toString())
}

func main(){
	test()
}
