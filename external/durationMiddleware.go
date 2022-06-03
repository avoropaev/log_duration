package external

import (
	"context"
	"errors"
	"time"
)

var ErrDeadline = errors.New("deadline reached")

type durationMiddleware struct {
	next Client
}

func DurationMiddleware(next Client) *durationMiddleware {
	return &durationMiddleware{next}
}

func (c *durationMiddleware) GetSomething(ctx context.Context, number int) (res int, err error) {
	start := time.Now()

	defer func() {
		if time.Since(start).Seconds() > 1 {
			err = ErrDeadline
		}
	}()

	return c.next.GetSomething(ctx, number)
}
