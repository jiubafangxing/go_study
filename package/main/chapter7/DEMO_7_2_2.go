	
package main
import(
	"log"
)
type MyError struct{

}

func (*MyError) Error()string{
	return "error"
}
func test1(a int )(string,error){
	if a < 0{
		var e *MyError = new(MyError)
		return "hello, err",e
	}else{
		return "success",nil 
	}
}

func test(){
	result, err := test1(1)	
	if(nil != err){
		log.Fatalln("err is not nil")
	}else{
		log.Println(result)
	}
}

func main(){
	test()
}
