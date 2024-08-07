	
package chapter4
import(
"log"
)
func DEMO_4_5_2test(){
	 x :=1
	 y :=2
	 defer func(a int){
		log.Println("x,y",a,y)		
	 }(x)
	 x +=100
	 y +=100
}

func DEMO_4_5_2main(){
	DEMO_4_5_2test()
}
