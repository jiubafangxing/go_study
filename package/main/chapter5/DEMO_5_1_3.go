	
package chapter5
import(
	"log"
)
func DEMO_5_1_3test(){
	s := `
		public class HelloWorld{
			public static void DEMO_5_1_3main(String[] args){
				System.out.println("hello world");
			}
		}

	`
	log.Println(s)
}

func DEMO_5_1_3main(){
	DEMO_5_1_3test()
}
