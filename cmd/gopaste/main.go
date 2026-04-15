package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {

	flag.Parse()

	s, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, s)
}
