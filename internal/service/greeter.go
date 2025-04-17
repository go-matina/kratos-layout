package service

import (
	"context"
	"fmt"

	pb "github.com/go-matina/kratos-layout/api/greeter"
	"github.com/go-matina/kratos-layout/internal/biz"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// Ping implements greeter.GreeterServer.
func (s *GreeterService) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &pb.PingReply{
		Message: fmt.Sprintf("Hello %s", g.Hello),
	}, nil
}
