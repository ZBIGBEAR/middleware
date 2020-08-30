package mw

import (
	"context"
	"fmt"
)

func LoggerMW(ctx context.Context, msg string, next MsgHandler) error {
	fmt.Println("LoggerMW before ")
	return next(ctx, msg)
}
