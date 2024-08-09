package chapter12

import (
	"log"
	"os"
	"testing"
)

func Add1214(a, b int) int {
	return a + b
}

func TestAdd1214(t *testing.T) {
	t.Parallel()
	if 3 != Add1214(1, 2) {
		t.FailNow()
	}
}

func ATestMain1214(m *testing.M) {
	log.Println("start test from TestMain")
	code := m.Run()
	log.Println("end test from TestMain")
	os.Exit(code)
}
