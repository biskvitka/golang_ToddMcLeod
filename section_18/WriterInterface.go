package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "hi")
	io.WriteString(os.Stdout, "hi again\n")
}