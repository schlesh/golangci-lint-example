package main

import (
	"fmt"
	"log"

	"github.com/containers/image/v5/transports/alltransports"
)

func main() {
	// Example usage: parse an image reference
	ref, err := alltransports.ParseImageName("docker://alpine:latest")
	if err != nil {
		log.Fatalf("Failed to parse image name: %v", err)
	}

	fmt.Println("Parsed image transport:", ref.Transport().Name())
}
