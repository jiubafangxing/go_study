package chapter7

type DEMO_7_4_1N int

type Stringer interface {
	toString() string
}

func (DEMO_7_4_1N) toString() string {
	return "a"
}

func init() {
	var _ Stringer = DEMO_7_4_1N(0)
}
func DEMO_7_4_1test() {

}

func DEMO_7_4_1main() {
	DEMO_7_4_1test()
}
