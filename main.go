package main

import (
	"context"
	"fmt"
	"log"

	"github.com/avoropaev/log_duration/external"
)

func main() {
	app := NewAppMiddleware(
		NewApp(
			external.NewClientDurationMiddleware(external.NewService1Client()),
		),
	)

	ctx := context.Background()

	result, err := app.Do(ctx, 123)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}
