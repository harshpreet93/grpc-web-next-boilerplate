package server

import (
	"errors"

	"context"
	"github.com/harshpreet93/lile_test"
)

func (s LileTestServer) Get(ctx context.Context, r *lile_test.GetRequest) (*lile_test.GetResponse, error) {
	return nil, errors.New("not yet implemented")
}
