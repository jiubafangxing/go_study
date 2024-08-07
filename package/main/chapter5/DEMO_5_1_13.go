	
package chapter5
import(
	"log"
	"unicode/utf8"
)
func DEMO_5_1_13test(){
	s := "你好世界"
	s = string(s[0:1]+s[3:4])
	log.Printf("string is %s ", s)
	log.Println(utf8.ValidString(s))
}

func DEMO_5_1_13main(){
	DEMO_5_1_13test()
}
