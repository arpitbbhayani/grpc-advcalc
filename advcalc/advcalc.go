package advcalc

import (
	"context"
	"math"
)

type AdvancedCalcServerLive struct {
}

func (*AdvancedCalcServerLive) Power(ctx context.Context, req *PowerRequest) (*PowerResponse, error) {
	return &PowerResponse{
		Pow: math.Pow(req.A, req.B),
	}, nil
}
