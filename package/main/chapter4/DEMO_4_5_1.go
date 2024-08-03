	
package main

import(
	"os"
	"log"
)
func main(){
	
	//f,err :=os.Open("./Template.go")
	f,err :=os.Open("./Template1.go")
	if nil != err{
		log.Fatalln(err)	
	}
	defer f.Close()
}
