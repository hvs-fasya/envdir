package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, os.Getenv("A_ENV")+" "+os.Getenv("B_ENV"))
}
