	
package main
import(
	"log"
)
func test(){
	s := `
		public class HelloWorld{
			public static void main(String[] args){
				System.out.println("hello world");
			}
		}

	`
	log.Println(s)
}

func main(){
	test()
}
