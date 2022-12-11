package main

import (
	"context"
	"fmt"

	"middleware/middleware"
	middlewarechain "middleware/middleware_chain"
	middlewarecontext "middleware/middleware_context"
)

func HandlerMsg(ctx context.Context, msg string) error {
	fmt.Println("HandlerMsg:", msg)
	return nil
}

func main() {
	// 方案一
	fmt.Println("===方案一 begin")
	m1 := middleware.NewMiddlewareManager(HandlerMsg)
	m1.Register(middleware.TimeCostMW, middleware.FilterMW, middleware.LoggerMW)
	if err := m1.Exec(context.Background(), "hello chain"); err != nil {
		panic(err)
	}
	fmt.Println("===方案一 end")

	// 方案二
	fmt.Println("===方案二 begin")
	m2 := middlewarecontext.NewMyContext()
	m2.Register(
		middlewarecontext.TimeCostMW,
		middlewarecontext.FilterMW,
		middlewarecontext.LoggerMW)
	if err := m2.Exec(); err != nil {
		panic(err)
	}
	fmt.Println("===方案二 end")

	// 方案三
	fmt.Println("===方案三 begin")
	ctx := context.Background()
	m3 := middlewarechain.TimeCostMW(ctx, func(ctx context.Context) error {
		PrintMsg("test")
		return nil
	})
	m4 := middlewarechain.FilterMW(ctx, m3)
	m5 := middlewarechain.LoggerMW(ctx, m4)
	if err := m5(ctx); err != nil {
		fmt.Println(err)
	}
	fmt.Println("===方案三 end")
}

func PrintMsg(msg string) {
	fmt.Println("PrintMsg:" + msg)
}

/*
===方案一 begin
TimeCost before
FinlterMW begin
LoggerMW before
HandlerMsg: hello chain
LoggerMW end
FinlterMW end
TimeCostMW:cost 1000428754
===方案一 end
===方案二 begin
TimeCost before
FinlterMW begin
LoggerMW before
LoggerMW end
FinlterMW end
TimeCostMW:cost 1000588399
===方案二 end
*/
