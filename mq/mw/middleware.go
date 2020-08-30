package mw

import "context"

// 消息处理函数
type MsgHandler func(ctx context.Context,
	msg string) error

// 插件类型
type MiddleWareFunc func(ctx context.Context,
	msg string, next MsgHandler) error
