package middlewarefor

import "context"

type Handler func(ctx context.Context) error

type Middleware func(next Handler) Handler

type MiddlewareManager struct {
	middlewares []Middleware
}

func NewMiddlewareManager(middlewares ...Middleware) *MiddlewareManager {
	return &MiddlewareManager{
		middlewares: middlewares,
	}
}

func (m *MiddlewareManager) Register(middlewares ...Middleware) {
	m.middlewares = append(m.middlewares, middlewares...)
}

func (m *MiddlewareManager) Exec(ctx context.Context) error {
	handler := defaultHandler
	for i := range m.middlewares {
		handler = m.middlewares[len(m.middlewares)-i-1](handler)
	}

	return handler(ctx)
}

func defaultHandler(ctx context.Context) error {
	return nil
}
