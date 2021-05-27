package main

import (
	"grpc-demo/domain/stream_operation"
	"log"
)

func main() {
	a := &stream_operation.Average{}

	dump(a)				// Error

	a.Feed(20)
	dump(a)				// 20

	a.Feed(30)
	dump(a)				// 25
}

func dump(a *stream_operation.Average) {
	average, err := a.Fetch()
	log.Printf("Average: %v, error: %v", average, err)
}
