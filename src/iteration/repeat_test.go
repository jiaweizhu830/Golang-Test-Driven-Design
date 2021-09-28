package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Run benchmark: go test -bench=.
// benchmark by default runs sequentially
// the framework will determine what is a "good" value for b.N to let you have some decent results.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	Repeat("a", 5)
	// Output: "aaaaa"
}
