package middlewarechain

import (
	"context"
	"fmt"
	"time"
)

func TimeCostMW(ctx context.Context, next Handler) Handler {
	return func(ctx context.Context) error {
		fmt.Println("TimeCost before")
		bTime := time.Now()
		err := next(ctx)
		fmt.Println(fmt.Sprintf("TimeCostMW:cost %d", time.Since(bTime)))
		return err
	}
}
