package main

import (
	"log"
	"os"

	"github.com/paketo-buildpacks/yarn-install/cmd/setup-symlinks/internal"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = internal.Run(os.Args[0], wd)
	if err != nil {
		log.Fatal(err)
	}
}
