	
package main
import(
	"os"
	"log"
	"os/signal"
	"sync"
	"syscall"
)
var exits = &struct{
	sync.RWMutex
	funcs []func()
	signals chan os.Signal
}{}
func atExit(f func()){
	exits.Lock()
	defer exits.Unlock()
	exits.funcs = append(exits.funcs, f)
}
func waitExit(){
	if nil == exits.signals{
		exits.signals = make(chan os.Signal)
		signal.Notify(exits.signals, syscall.SIGINT, syscall.SIGTERM)
	}
	exits.RLock()
	for _, f := range exits.funcs{
		defer f()
	}
	exits.RUnlock()
	<- exits.signals
}

func test(){
	atExit(func(){log.Println("exit1")})
	atExit(func(){log.Println("exit2")})
	waitExit()
}

func main(){
	test()
}
