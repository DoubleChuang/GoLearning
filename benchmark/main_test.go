package main

import (
	"testing"
)

var num = 11110
var str_num = "11110"

func TestPrint01(t *testing.T) {
	if print01(num) != str_num {
		t.Errorf("print(%d) != %s\n", num, print01(num))
	}
}
func TestPrint02(t *testing.T) {
	if print02(int64(num)) != str_num {
		t.Errorf("print(%d) != %s\n", num, print02(int64(num)))
	}
}
func TestPrint03(t *testing.T) {
	if print03(num) != str_num {
		t.Errorf("print(%d) != %s\n", num, print03(num))
	}
}
func BenchmarkPrint01(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		print01(11110)
	}
}
func BenchmarkPrint02(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		print02(11110)
	}
}
func BenchmarkPrint03(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		print01(11110)
	}
}
