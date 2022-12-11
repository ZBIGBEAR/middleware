package middleware

import (
	"context"
	"fmt"
)

func LoggerMW(ctx context.Context, msg string, next Handler) error {
	fmt.Println("LoggerMW before ")
	err := next(ctx, msg)
	fmt.Println("LoggerMW end ")
	return err
}
