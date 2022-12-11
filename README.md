# 插件的实现方式
经常看到一些框架有插件，具体插件怎么实现的呢？这几天review同事代码的时候，发现他写了一段从消息队列中获取并处理消息的代码，其中用到了插件。真正保存消息之前的一些操作，比如耗时统计、参数处理、日志处理、异常捕获等操作都以插件的形式实现。这段代码让我学习到了插件实现的原理。方案一是对这段代码的实现。

这里还总结了一下gin框架的中间件实现方式，是一种顺序实现，它把中间件封装到自定义的Context结构体中，通过```ctx.Next()```实现链式调用。方案二是对这段代码的实现。

还有一种链式调用中间件，它的入参是next，返回值也是一个next，实现链式调用。这种函数套函数的用法不是很美观，同时也不具备什么可读性。方案三是这个方案的实现。

# 示例

```
package main

import (
	"context"
	"fmt"

	"middleware/middleware"
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
```

结果
```
===方案一 begin
TimeCost before
FinlterMW begin
LoggerMW before 
HandlerMsg: hello chain
LoggerMW end 
FinlterMW end
TimeCostMW:cost 1000284305
===方案一 end
===方案二 begin
TimeCost before
FinlterMW begin
LoggerMW before 
LoggerMW end 
FinlterMW end
TimeCostMW:cost 1000699792
===方案二 end
===方案三 begin
LoggerMW before 
FinlterMW begin
TimeCost before
PrintMsg:test
TimeCostMW:cost 6045
FinlterMW end
LoggerMW end 
===方案三 end
```
