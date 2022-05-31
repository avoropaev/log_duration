package external

import (
	"context"
	"time"
)

type Service1Client interface {
	GetSomething(ctx context.Context, number int) (int, error)
}

type client struct {}

func NewService1Client() Service1Client {
	return client{}
}

func (c client) GetSomething(_ context.Context, number int) (int, error) {
	// если есть эта строка, то лог не должен писаться
	time.Sleep(time.Millisecond * 1001)

	return number * 2, nil
}
