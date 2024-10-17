package iteration

import (
	"fmt"
	"strings"
	"testing"
)

// Repeat test, example and benchmark
func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("g", 5)
	fmt.Println(repeated)
	// Output: ggggg
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// Go standard library repeat function test
func TestRepeatGoStd(t *testing.T) {
	repeated := strings.Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
