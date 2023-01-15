package middlewarefor

import (
	"context"
	"time"

	"log"
)

func TimeCostMW(next Handler) Handler {
	return func(ctx context.Context) error {
		begin := time.Now()
		err := next(ctx)
		log.Printf("[TimeCostMW] cost:%fs", time.Since(begin).Seconds())
		return err
	}
}
