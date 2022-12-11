package middlewarechain

import (
	"context"
	"fmt"
)

func LoggerMW(ctx context.Context, next Handler) Handler {
	return func(ctx context.Context) error {
		fmt.Println("LoggerMW before ")
		err := next(ctx)
		fmt.Println("LoggerMW end ")
		return err
	}
}
