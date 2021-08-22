package main

import (
	"context"
	"fmt"
	"time"
)

type processFn func(context.Context, *cmd) error

type cmd struct {
	name string
	req  interface{}
	res  interface{}
}

type ctxStartTimeKey struct{}

type Interceptor func(oldProcessFn processFn) (newProcessFn processFn)

func InterceptorChain(interceptors ...Interceptor) Interceptor {
	return func(p processFn) processFn {
		chain := p
		for i := len(interceptors) - 1; i >= 0; i-- {
			chain = buildInterceptor(interceptors[i], chain)
		}
		return chain
	}
}

func buildInterceptor(interceptor Interceptor, oldProcess processFn) processFn {
	return interceptor(oldProcess)
}

func fixedInterceptor() Interceptor {
	return func(next processFn) processFn {
		return func(ctx context.Context, cmd *cmd) error {
			start := time.Now()
			fmt.Printf("ctx1--------------->"+"%+v\n", ctx)
			fmt.Printf("cmd fixedInterceptor1 start--------------->"+"%+v\n", cmd)
			ctx = context.WithValue(ctx, ctxStartTimeKey{}, start)

			err := next(ctx, cmd)
			fmt.Printf("cmd fixedInterceptor1 end--------------->"+"%+v\n", cmd)
			return err
		}
	}
}

func TestInterceptor() Interceptor {
	return func(next processFn) processFn {
		return func(ctx context.Context, cmd *cmd) error {
			start := time.Now()
			fmt.Printf("ctx2--------------->"+"%+v\n", ctx)
			fmt.Printf("cmd fixedInterceptor222 start--------------->"+"%+v\n", cmd)
			ctx = context.WithValue(ctx, ctxStartTimeKey{}, start)

			err := next(ctx, cmd)
			fmt.Printf("cmd fixedInterceptor22 end--------------->"+"%+v\n", cmd)
			return err
		}
	}
}

func main() {
	chain := InterceptorChain(fixedInterceptor(), TestInterceptor())

	ctx := context.Background()
	ctx = context.WithValue(ctx, "hello", "world")
	_ = chain(func(ctx context.Context, c *cmd) error {
		fmt.Println("hello world")
		return nil
	})(ctx, &cmd{})
}
