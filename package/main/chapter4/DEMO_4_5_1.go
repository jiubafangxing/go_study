	
package chapter4

import(
	"os"
	"log"
)
func DEMO_4_5_1main(){
	
	//f,err :=os.Open("./Template.go")
	f,err :=os.Open("./Template1.go")
	if nil != err{
		log.Fatalln(err)	
	}
	defer f.Close()
}
