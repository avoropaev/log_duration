package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/avoropaev/log_duration/external"
)

// мейн это типо мидлварь сервера
func main() {
	app := NewApp(
		external.NewClientDurationMiddleware(external.NewService1Client()),
	)

	ctx := context.Background()

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

	result, err := app.Do(ctx, 123)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}