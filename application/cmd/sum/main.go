package main

import (
	"grpc-demo/domain/simple_operation"
	"log"
)

func main() {
	simpleOperation := simple_operation.SumOperation{}
	result := simpleOperation.Execute(6, 9);

	log.Printf("The result is: %v", result)
}
