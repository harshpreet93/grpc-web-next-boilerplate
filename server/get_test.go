package server

import (
	"testing"

	"context"
	"github.com/harshpreet93/lile_test"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ctx := context.Background()
	req := &lile_test.GetRequest{}

	res, err := cli.Get(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
