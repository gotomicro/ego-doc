package main

import (
	"context"
	"fmt"
	"time"
)

type processFn func(*cmd) error

type cmd struct {
	name string
	req  interface{}
	res  interface{}
}

type ctxStartTimeKey struct{}

type Interceptor func(ctx context.Context, oldProcessFn processFn) (newProcessFn processFn)

func InterceptorChain(interceptors ...Interceptor) Interceptor {
	return func(ctx context.Context, p processFn) processFn {
		chain := p
		for i := len(interceptors) - 1; i >= 0; i-- {
			fmt.Printf("222222--------------->"+"%+v\n", ctx)
			chain = buildInterceptor(interceptors[i], ctx, chain)
			fmt.Printf("33333--------------->"+"%+v\n", ctx)
		}
		return chain
	}
}

func buildInterceptor(interceptor Interceptor, ctx context.Context, oldProcess processFn) processFn {
	return interceptor(ctx, oldProcess)
}

func fixedInterceptor() Interceptor {
	return func(ctx context.Context, next processFn) processFn {
		return func(cmd *cmd) error {
			start := time.Now()
			fmt.Printf("ctx--------------->"+"%+v\n", ctx)
			fmt.Printf("cmd fixedInterceptor start--------------->"+"%+v\n", cmd)
			ctx = context.WithValue(ctx, ctxStartTimeKey{}, start)

			err := next(cmd)
			fmt.Printf("cmd fixedInterceptor end--------------->"+"%+v\n", cmd)
			return err
		}
	}
}

func TestInterceptor() Interceptor {
	return func(ctx context.Context, next processFn) processFn {
		return func(cmd *cmd) error {
			start := time.Now()
			fmt.Printf("ctx--------------->"+"%+v\n", ctx)
			fmt.Printf("cmd fixedInterceptor222 start--------------->"+"%+v\n", cmd)
			ctx = context.WithValue(ctx, ctxStartTimeKey{}, start)

			err := next(cmd)
			fmt.Printf("cmd fixedInterceptor22 end--------------->"+"%+v\n", cmd)
			return err
		}
	}
}

func main() {
	chain := InterceptorChain(fixedInterceptor(), TestInterceptor())

	ctx := context.Background()
	ctx = context.WithValue(ctx, "hello", "world")
	_ = chain(ctx, func(c *cmd) error {
		fmt.Println("hello world")
		return nil
	})(&cmd{})
}
