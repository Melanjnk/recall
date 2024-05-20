package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s := fmt.Sprintf("%s", "123")
	io.WriteString(os.Stdout, s)
}
