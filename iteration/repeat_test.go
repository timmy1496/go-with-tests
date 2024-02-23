package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func TestRepeat(t *testing.T) {
	repeatCount := 10
	repeated := Repeat("a", repeatCount)
	var expected string

	for i := 0; i < repeatCount; i++ {
		expected += "a"
	}

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
