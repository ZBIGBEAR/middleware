package main

import (
	"context"
	"fmt"
	"middleware/mq"
	. "middleware/mq/mw"
)

func HandlerMsg(ctx context.Context, msg string) error {
	fmt.Println("HandlerMsg:", msg)
	return nil
}

func main() {
	mq := mq.NewMQHandler(HandlerMsg)
	mq.Register(TimeCostMW, FilterMW, LoggerMW)
	mq.Exec(context.Background(), "hello chain")
}

//输出：
//TimeCost before
//FinlterMW
//LoggerMW before
//HandlerMsg: hello chain
//TimeCostMW:cost 1000065900
