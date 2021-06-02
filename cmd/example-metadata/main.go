package main

import (
	"fmt"
	"log"

	"github.com/spiretechnology/spireav"
)

func main() {

	filename := "./cmd/example-metadata/sample.mp4"

	// Open the file
	ctx, err := spireav.OpenFileContext(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer ctx.Close()

	// Parse the metadata
	meta, err := ctx.Metadata()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Loop through the keys and values
	for key, value := range meta {
		fmt.Println(key + " => " + value)
	}

}
