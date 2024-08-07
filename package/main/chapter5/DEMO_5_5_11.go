	
package chapter5
import(
	"log"
)
func DEMO_5_5_11test(){
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

func DEMO_5_5_11main(){
	DEMO_5_5_11test()
}
