	
package main
import(
	"log"
)

type node struct{
	_ int
	id int
	next *node
}
func test(){
	var n1 node =node{
		id:1,
	}
	n2 := node{
		id:2,
		next: &n1,
	}
	log.Printf("the node is %v", n1)
	log.Printf("the node is %v", n2)


}

func main(){
	test()
}
