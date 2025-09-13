package main

import (
    "os"
    "github.com/brettfirecore/hello/hello"
)

func main() {
    hello.Print(os.Stdout)
}
