package main

import (
	"fmt"
	"github.com/tyhal/pinup/upgrade"
	"log"
	"os"
)

// TODO use cobra lib for CLI

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: pinup <Dockerfile>")
	}

	filename := os.Args[1]
	fmt.Println("Opening " + filename)

	fileIn, err := os.Open(filename)

	if err != nil {
		log.Fatal("Couldnt open " + filename)
	}
	defer fileIn.Close()

	upgrade.Docker(fileIn)
}
