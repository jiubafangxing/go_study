	
package chapter7
import(
	"log"
)
type MyError struct{

}

func (*MyError) Error()string{
	return "error"
}
func DEMO_7_2_2test1(a int )(string,error){
	if a < 0{
		var e *MyError = new(MyError)
		return "hello, err",e
	}else{
		return "success",nil 
	}
}

func DEMO_7_2_2test(){
	result, err := DEMO_7_2_2test1(1)	
	if(nil != err){
		log.Fatalln("err is not nil")
	}else{
		log.Println(result)
	}
}

func DEMO_7_2_2main(){
	DEMO_7_2_2test()
}
