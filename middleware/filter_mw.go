package middleware

import (
	"context"
	"fmt"
	"time"
)

func FilterMW(ctx context.Context, msg string, next Handler) error {
	fmt.Println("FinlterMW begin")
	time.Sleep(time.Second)
	err := next(ctx, msg)
	fmt.Println("FinlterMW end")
	return err
}
