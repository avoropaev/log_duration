package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/avoropaev/log_duration/external"
)

type App interface {
	Do(ctx context.Context, number int) (int, error)
}

type app struct {
	svc external.Client
}

func NewApp(svc external.Client) App {
	return app{svc}
}

func (a app) Do(ctx context.Context, number int) (int, error) {
	resp, err := a.svc.GetSomething(ctx, number)
	if err != nil {
		if errors.Is(err, external.ErrDeadline) {
			fmt.Println("slow request")
			err = nil
		}
	}
	return resp, err
}
