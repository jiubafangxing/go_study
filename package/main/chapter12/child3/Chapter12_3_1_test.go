package child3

import (
	"testing"
)

func TestC1231Add(t *testing.T) {
	if C1231Add(1, 2) != 0 {
		t.Fatal()
	} else if C1231Add(2, 3) != 5 {
		t.Fatal()
	}

}
