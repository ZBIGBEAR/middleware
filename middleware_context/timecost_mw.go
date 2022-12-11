package middlewarecontext

import (
	"fmt"
	"time"
)

func TimeCostMW(ctx *MyContext) error {
	fmt.Println("TimeCost before")
	bTime := time.Now()
	if err := ctx.Next(); err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("TimeCostMW:cost %d", time.Since(bTime)))
	return nil
}
