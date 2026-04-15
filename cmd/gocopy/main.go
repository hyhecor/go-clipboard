package main

import (
	"flag"
	"io"
	"os"

	"github.com/atotto/clipboard"
)

func main() {

	flag.Parse()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	s := string(data)

	err = clipboard.WriteAll(s)
	if err != nil {
		panic(err)
	}
}
