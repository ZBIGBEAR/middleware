package middlewarecontext

type MiddleWareFunc func(ctx *MyContext) error

type MyContext struct {
	middlewares []MiddleWareFunc
	idx         int
	maxIdx      int
}

func NewMyContext() *MyContext {
	return &MyContext{
		middlewares: make([]MiddleWareFunc, 0),
	}
}

// 执行下一个middleware
func (m *MyContext) Next() error {
	if m.idx < m.maxIdx-1 {
		m.idx += 1
		return m.middlewares[m.idx](m)
	}

	return nil
}

// 终止middleware
func (m *MyContext) Abort() {
	m.idx = m.maxIdx
}

func (m *MyContext) Register(middlewares ...MiddleWareFunc) {
	m.middlewares = append(m.middlewares, middlewares...)
	m.maxIdx = len(m.middlewares)
}

func (m *MyContext) Exec() error {
	// 从第一个middleware开始执行
	return m.middlewares[0](m)
}
