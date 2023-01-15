package middlewarefor

import (
	"context"

	"log"
)

func LoggerMW(next Handler) Handler {
	return func(ctx context.Context) error {
		log.Printf("[LoggerMW] befor")
		err := next(ctx)
		log.Printf("[LoggerMW] end")
		return err
	}
}
