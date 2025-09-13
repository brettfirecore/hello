package hello

import (
	"bytes"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	Print(&buf)

	got := buf.String()
	want := "Hello, world!\n"
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

// Examples must print to stdout for the `// Output:` check to work.
func ExamplePrint() {
	Print(os.Stdout)
	// Output:
	// Hello, world!
}
