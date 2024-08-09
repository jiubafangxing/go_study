package chapter12

import (
	"testing"
)

func Add1215(a, b int) int {
	return a + b
}

func Minus1215(a, b int) int {
	return a - b
}

func TestAdd1215(t *testing.T) {
	t.Parallel()
	if 3 != Add1215(1, 2) {
		t.FailNow()
	}
}

func TestMinus1215(t *testing.T) {
	t.Parallel()
	if 3 != Minus1215(4, 1) {
		t.FailNow()
	}
}

func TestMain(m *testing.M) {
	//match := func(pat, fname string) bool {
	//	//default  test all method
	//	return true
	//}
	//testCase := []testing.InternalTest{
	//	{"a", TestAdd1215},
	//	{"b", TestMinus1215},
	//}
	//benchmarks := []testing.InternalBenchmark{}
	//examples := []testing.InternalExample{}
	//fuzzTargets := []testing.InternalFuzzTarget{}
	//m = testing.MainStart(nil, testCase, benchmarks, fuzzTargets, examples)
	//os.Exit(m.Run())
}
