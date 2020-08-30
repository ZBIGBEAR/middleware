package mq

import (
	"context"
	. "middleware/mq/mw"
)

type MQHandler struct {
	handler     MsgHandler
	middlewares []MiddleWareFunc
}

func NewMQHandler(handler MsgHandler) *MQHandler {
	return &MQHandler{
		handler: handler,
	}
}

func (m *MQHandler) Register(middlewares ...MiddleWareFunc) {
	m.middlewares = append(m.middlewares, middlewares...)
}

func (m *MQHandler) Exec(ctx context.Context, msg string) error {
	handlerFunc := func(ctx context.Context, msg string, next MsgHandler) error {
		return m.handler(ctx, msg)
	}
	m.middlewares = append(m.middlewares, handlerFunc)
	callChain := m.mkCallChain(m.middlewares)
	return callChain(ctx, msg)
}

func (m *MQHandler) mkCallChain(
	middlewares []MiddleWareFunc) MsgHandler {
	if len(middlewares) <= 0 {
		return nil
	}
	middleware := middlewares[0]
	handler := m.mkCallChain(middlewares[1:])
	return func(ctx context.Context, msg string) error {
		return middleware(ctx, msg, handler)
	}
}
