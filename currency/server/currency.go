package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/nicholasjackson/building-microservices-youtube/currency/protos/currency"
)

type Currency struct {
	logger hclog.Logger
}

func NewCurrency(log hclog.Logger) *Currency {
	return &Currency{log}
}

func (c *Currency) GetRate(context context.Context, rateRequest *protos.RateRequest) (*protos.RateResponse, error) {
	c.logger.Info("Handle GetRate", "base", rateRequest.GetBase(), "destination", rateRequest.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
