package arithmetic_server

import (
	"context"
	"saranshsaini/arithmetic_pb"
)


type ArithmeticServer struct {
	arithmetic_pb.UnimplementedArithmeticServiceServer
}

func (s *ArithmeticServer) Add(ctx context.Context, req *arithmetic_pb.AddRequest) (*arithmetic_pb.AddResponse, error) {
	return &arithmetic_pb.AddResponse{
		C: req.A + req.B,
	}, nil
}

func (s *ArithmeticServer) Subtract(ctx context.Context, req *arithmetic_pb.SubtractRequest) (*arithmetic_pb.SubtractResponse, error) {
	return &arithmetic_pb.SubtractResponse{
		C: req.A - req.B,
	}, nil
}