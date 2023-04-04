package main

// testのpackage
import "testing"

// debugをするかどうか切り替える用の変数
var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("skip")
	}

	v := IsOne(i)

	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}