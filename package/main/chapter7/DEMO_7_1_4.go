	
package main
import(
	"log"
)
type tester interface{
	stringer
	test()
}
type stringer interface{
	toString()(string)
}
type N int

func (n *N) test(){
	log.Println(*n)
}

func (n *N)toString()(string){
	return "a"
}

func test1(){
	var n N = 1
	var tester1 tester = &n
	tester1.test()
	log.Println(tester1.toString())
	var str1 stringer = tester1
	log.Println(str1.toString())
	//can not be used on  this way
	//var tester2 tester = str1
	//tester2.test()
}

func main(){
	test1()
}
