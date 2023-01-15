package middlewarefor

import (
	"context"
	"fmt"
	"log"
	"runtime"
)

func RecoveryMW(next Handler) Handler {
	return func(ctx context.Context) error {
		log.Printf("[RecoveryMW] befor")
		defer func() {
			log.Printf("[RecoveryMW] end")
			if r := recover(); r != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				log.Printf("err:%v, buf:%s", err, string(buf))
			}
		}()

		return next(ctx)
	}
}
