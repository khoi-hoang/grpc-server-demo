package grpc_stream_operation

import (
	"grpc-demo/domain/stream_operation"
	"io"
	"log"
)

type Server struct {}

func (Server) AverageClientStream(server StreamOperation_AverageClientStreamServer) error {
	a := stream_operation.Average{}

	for {
		req, err := server.Recv()
		log.Printf("Request received: %v, error: %v", req, err)

		if err == io.EOF {
			result, err := a.Fetch()

			if err != nil {
				return err
			}

			return server.SendAndClose(&OperationResponse{Result: result})
		}

		if err != nil {
			return err
		}

		a.Feed(req.Num)
	}
}

func (Server) PrimeDecompose(request *OperationRequest, server StreamOperation_PrimeDecomposeServer) error {
	d := stream_operation.NewPrimeDecomposer()

	d.Feed(request.Num)

	for {
		result, err := d.Fetch()
		log.Printf("Dispatching result: %v, error: %v", result, err)

		if err == io.EOF {
			break
		}

		if err != nil {
			// Something went wrong in the domain
			return err
		}

		if err := server.Send(&OperationResponse{Result: result}); err != nil {
			// Server failed to dispatch the response
			return err
		}
	}

	return nil
}

func (s Server) AverageBidirectionalStream(server StreamOperation_AverageBidirectionalStreamServer) error {
	a := stream_operation.Average{}

	for {
		req, err := server.Recv()
		log.Printf("Received: %v, error: %v", req, err)

		if err == io.EOF {
			break
		}

		if err != nil {
			// Failed to receive from client
			return err
		}

		a.Feed(req.Num)
		result, err := a.Fetch()

		if err != nil {
			// Domain's error
			return err
		}

		log.Printf("Dispatching: %v", result)
		if err := server.Send(&OperationResponse{Result: result}); err != nil {
			// Failed to dispatch the response
			return err
		}
	}

	return nil
}

func (Server) mustEmbedUnimplementedStreamOperationServer() {
	panic("implement me")
}
