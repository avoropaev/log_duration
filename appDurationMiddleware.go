package main

import (
	"context"
	"log"
	"time"

	"github.com/avoropaev/log_duration/external"
)

type appDurationMiddleware struct {
	originalApp App
}

func NewAppMiddleware(app App) App {
	return appDurationMiddleware{
		originalApp: app,
	}
}

func (a appDurationMiddleware) Do(ctx context.Context, number int) (int, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)

		service1Duration, ok := ctx.Value(external.Service1QueryDurationContextKey).(time.Time)
		if ok && service1Duration.Second() > 1 {
			// если был запрос во внешний сервис, который длился больше 1 секунды, то "слишком долго" не логируем
			return
		}

		if duration.Seconds() > 2 {
			log.Print("слишком долго")
		}
	}()

	return a.originalApp.Do(ctx, number)
}
