package service

import (
	"context"
	"fmt"

	"github.com/alexadastra/ramme-template/pkg/api"
)

// RammeTemplate handles stuff
type RammeTemplate struct{}

// NewRammeTemplate creates new server
func NewRammeTemplate() api.RammeTemplateServer {
	return &RammeTemplate{}
}

// Ping returns "pong" if ping in pinged
func (rt *RammeTemplate) Ping(ctx context.Context, request *api.PingRequest) (*api.PingResponse, error) {
	if request.Value == "ping" {
		return &api.PingResponse{
			Code:  200,
			Value: "pong",
		}, nil
	}
	return nil, fmt.Errorf("unknown request message: %s", request.Value)
}
