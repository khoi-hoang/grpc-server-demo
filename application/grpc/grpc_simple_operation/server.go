package grpc_simple_operation

import (
	context "context"
	"grpc-demo/domain/simple_operation"
)

type Server struct{}

func (Server) Sum(ctx context.Context, request *SimpleOperationRequest) (*SimpleOperationResponse, error) {
	s := simple_operation.SumOperation{}

	first, second := request.First, request.Second
	result := s.Execute(first, second)

	return &SimpleOperationResponse{Result: result}, nil
}

func (Server) mustEmbedUnimplementedSimpleOperationServer() {
	panic("implement me")
}
