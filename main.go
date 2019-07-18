package main

import (
	"log"

	"github.com/hvs-fasya/envdir/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
