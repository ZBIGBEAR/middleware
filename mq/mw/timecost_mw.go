package mw

import (
	"context"
	"fmt"
	"time"
)

func TimeCostMW(ctx context.Context, msg string, next MsgHandler) error {
	fmt.Println("TimeCost before")
	bTime := time.Now()
	err := next(ctx, msg)
	fmt.Println(fmt.Sprintf("TimeCostMW:cost %d", time.Since(bTime)))
	return err
}
