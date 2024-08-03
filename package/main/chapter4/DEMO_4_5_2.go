	
package main
import(
"log"
)
func test(){
	 x :=1
	 y :=2
	 defer func(a int){
		log.Println("x,y",a,y)		
	 }(x)
	 x +=100
	 y +=100
}

func main(){
	test()
}
