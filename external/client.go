package external

import (
	"context"
	"time"
)

type Client interface {
	GetSomething(ctx context.Context, number int) (int, error)
}

type clientImpl struct{}

func NewClient() Client {
	return &clientImpl{}
}

func (c *clientImpl) GetSomething(ctx context.Context, number int) (int, error) {
	time.Sleep(1 * time.Second)
	return number * 2, nil
}
