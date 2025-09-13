package hello

import (
	"fmt"
	"io"
)

func Print(w io.Writer) {
	fmt.Fprintln(w, "Hello, world!")
}
