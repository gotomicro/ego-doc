package gracefulstop

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 普通server stop
func CommonStop() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "127.0.0.1:0",
		Handler: mux,
	}

	go func() {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-sig:
				server.Close()
			}
		}
	}()
	server.ListenAndServe()
}

// 正常的关闭
func GracefulStop() {
	g, ctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "127.0.0.1:0",
		Handler: mux,
	}

	g.Go(func() error {
		<-ctx.Done()
		fmt.Println("http ctx done")
		return server.Shutdown(context.TODO())
	})
	// http server
	g.Go(func() error {
		fmt.Println("http")
		return server.ListenAndServe()
	})
	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case signalInfo := <-sig:
				// do something
				return fmt.Errorf("signal info %v", signalInfo)
			}
		}
	})

	err := g.Wait() // first error return
	fmt.Println(err)
}
