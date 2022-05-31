package external

import (
	"context"
	"time"
)

type contextKey string

const Service1QueryDurationContextKey = contextKey("service1_query_duration")

type clientDurationMiddleware struct {
	originalClient Service1Client
}

func NewClientDurationMiddleware(originalClient Service1Client) Service1Client {
	return clientDurationMiddleware{
		originalClient: originalClient,
	}
}

func (c clientDurationMiddleware) GetSomething(ctx context.Context, number int) (int, error) {
	start := time.Now()
	defer func() {
		ctx = context.WithValue(ctx, Service1QueryDurationContextKey, time.Since(start))
	}()

	return c.originalClient.GetSomething(ctx, number)
}
