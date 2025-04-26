package rfms

import (
	"context"
	"fmt"
)

type ListVehiclesRequest struct{}

type ListVehiclesResponse struct{}

func (c *Client) ListVehicles(ctx context.Context, request *ListVehiclesRequest) (*ListVehiclesResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
