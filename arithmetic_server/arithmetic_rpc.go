package arithmetic_server

import (
	"context"
	"saranshsaini/arithmetic_pb"
)


type SampleServer struct {
	arithmetic_pb.UnimplementedArithmeticServiceServer
}

func (s *SampleServer) Add(ctx context.Context, req *arithmetic_pb.AddRequest) (*arithmetic_pb.AddResponse, error) {
	return &arithmetic_pb.AddResponse{
		C: req.A + req.B,
	}, nil
}