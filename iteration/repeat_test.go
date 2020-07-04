package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 4)
	expected := "aaaa"

	if actual != expected {
		t.Errorf("expected %q but actual %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	output := Repeat("a", 3)
	fmt.Println(output)
	// Output: aaa
}
