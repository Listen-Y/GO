package study_testing

import (
	"testing"
)

func TestAdd(t *testing.T) {
	res := add(10, 20)
	if res != 30 {
		t.Fatal("test add 不符合预期")
	}
	t.Logf("test add 符合预期 res: %v\n", res)
}
