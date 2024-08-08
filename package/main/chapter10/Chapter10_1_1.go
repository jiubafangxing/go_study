package chapter10

type c1011data struct {
	x int
	y string
}

type C1011data struct{
	x int `field:"x"`
	y string
}

func BuildC1011Data() *c1011data {
	return &c1011data{
		x: 1,
	}
}
