package middlewarechain

import (
	"context"
	"fmt"
	"time"
)

func FilterMW(ctx context.Context, next Handler) Handler {
	return func(ctx context.Context) error {
		fmt.Println("FinlterMW begin")
		time.Sleep(time.Second)
		err := next(ctx)
		fmt.Println("FinlterMW end")
		return err
	}
}
