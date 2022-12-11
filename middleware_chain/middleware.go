package middlewarechain

import "context"

type Handler func(ctx context.Context) error

type MiddleWareFunc func(ctx context.Context, next Handler) Handler
