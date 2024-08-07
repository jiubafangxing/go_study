	
package chapter5
import(
	"log"
)
func DEMO_5_5_3test(){
	u := struct{
		id int
		name string

	}{

		id:1,
		name:"laoli",
	}
	log.Println(u)
	type file struct{
		name string
		attr  struct{
			owner int
			perm int
		}
	}

	f1 := file{

		name:"xxx.pdf",
	}
	f1.attr.owner = 1
	f1.attr.perm = 2
	log.Println(f1)

}

func DEMO_5_5_3main(){
	DEMO_5_5_3test()
}
