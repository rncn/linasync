package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagVersion = flag.Bool("version", false, "Show Version")
)

func main() {
	flag.Parse()

	if *flagVersion {
		showVersionToStdOut()
	}
}

func showVersionToStdOut() {
	fmt.Println("0.0.1")

	os.Exit(0)
}
