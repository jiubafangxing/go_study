	
package chapter5
import(
	"log"
	"unicode/utf8"
)
func DEMO_5_1_14test(){
	s := "你好世界a"
	log.Println(utf8.RuneCountInString(s))


}

func DEMO_5_1_14main(){
	DEMO_5_1_14test()
}
