package middleware

import "context"

// 处理函数
type Handler func(ctx context.Context,
	msg string) error

// 插件类型
type MiddleWareFunc func(ctx context.Context,
	msg string, next Handler) error

type MiddlewareManager struct {
	handler     Handler
	middlewares []MiddleWareFunc
}

func NewMiddlewareManager(handler Handler) *MiddlewareManager {
	return &MiddlewareManager{
		handler: handler,
	}
}

func (m *MiddlewareManager) Register(middlewares ...MiddleWareFunc) {
	m.middlewares = append(m.middlewares, middlewares...)
}

func (m *MiddlewareManager) Exec(ctx context.Context, msg string) error {
	handlerFunc := func(ctx context.Context, msg string, next Handler) error {
		return m.handler(ctx, msg)
	}
	m.middlewares = append(m.middlewares, handlerFunc)

	callChain := m.mkCallChain(m.middlewares)
	return callChain(ctx, msg)
}

func (m *MiddlewareManager) mkCallChain(
	middlewares []MiddleWareFunc) Handler {
	if len(middlewares) <= 0 {
		return nil
	}

	return func(ctx context.Context, msg string) error {
		return middlewares[0](ctx, msg, m.mkCallChain(middlewares[1:]))
	}
}
