package main

import (
	"context"
	"fmt"
	"github.com/avoropaev/log_duration/external"
)

func main() {
	app := NewApp(external.DurationMiddleware(external.NewClient()))

	result, err := app.Do(context.Background(), 123)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
