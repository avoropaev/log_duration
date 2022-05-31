package main

import (
	"context"
	"time"

	"github.com/avoropaev/log_duration/external"
)

type App interface {
	Do(ctx context.Context, number int) (int, error)
}

type app struct {
	service1 external.Service1Client
}

func NewApp(service1 external.Service1Client) App {
	return app{
		service1: service1,
	}
}

func (a app) Do(ctx context.Context, number int) (int, error) {
	time.Sleep(time.Millisecond * 2001)

	return a.service1.GetSomething(ctx, number)
}