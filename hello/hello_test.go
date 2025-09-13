package hello_test

import (
	"bytes"
	"testing"

	"github.com/brettfirecore/hello/hello"
)

func TestPrintPrintsHelloMessageToTerminal(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	hello.Print(&buf)

	if got, want := buf.String(), "Hello, world!\n"; got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}
