package middlewarecontext

import (
	"fmt"
)

func LoggerMW(ctx *MyContext) error {
	fmt.Println("LoggerMW before ")
	if err := ctx.Next(); err != nil {
		return err
	}
	fmt.Println("LoggerMW end ")
	return nil
}
