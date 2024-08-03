	
package main
import(
	"log"
	"unicode/utf8"
)
func test(){
	s := "你好世界a"
	log.Println(utf8.RuneCountInString(s))


}

func main(){
	test()
}
