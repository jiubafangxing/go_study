	
package main
import(
	"log"
)
func test(){
	type attr  struct{
		perm int

	}
	type File struct{
		name string
		attr 
	}
	file := File{
		name:"wang",
		attr:attr{
			perm:1,
		},
	}
	log.Println(file)
	log.Println(file.perm)
}

func main(){
	test()
}
