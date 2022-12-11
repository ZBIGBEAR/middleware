package middlewarecontext

import (
	"fmt"
	"time"
)

func FilterMW(ctx *MyContext) error {
	fmt.Println("FinlterMW begin")
	time.Sleep(time.Second)
	if err := ctx.Next(); err != nil {
		return err
	}
	fmt.Println("FinlterMW end")
	return nil
}
