package mw

import (
	"context"
	"fmt"
	"time"
)

func FilterMW(ctx context.Context, msg string, next MsgHandler) error {
	fmt.Println("FinlterMW")
	time.Sleep(time.Second)
	return next(ctx, msg)
}
