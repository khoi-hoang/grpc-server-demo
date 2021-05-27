package main

import (
	"grpc-demo/domain/stream_operation"
	"io"
	"log"
)

func main() {
	primeDecomposer := stream_operation.NewPrimeDecomposer()
	primeDecomposer.Feed(237868273451)		// How about 8247592349359818797

	for {
		prime, err := primeDecomposer.Fetch()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not decompose: %v", err)
		}

		log.Printf("Next prime number: %v", prime)
	}
}