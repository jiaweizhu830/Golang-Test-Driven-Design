package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		// %d placeholder for integer
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Examples are compiled (and optionally executed) as part of a package's test suite.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
