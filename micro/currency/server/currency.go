package server

import (
	"context"

	"github.com/alexbarksdale/GoGo/micro/currency/pb"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
	pb.UnimplementedCurrencyServer
}

func NewCurrencyServer(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, req *pb.RateRequest) (*pb.RateResponse, error) {
	c.log.Info("Handle GetRate", "base:", req.GetBase(), "destination:", req.GetDestination())

	return &pb.RateResponse{Rate: 0.5}, nil
}
