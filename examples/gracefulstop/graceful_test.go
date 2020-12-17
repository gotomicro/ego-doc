package gracefulstop

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestGracefulStop(t *testing.T) {
	go func() {
		time.Sleep(5 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	}()
	GracefulStop()
}

func TestCommonShutdown(t *testing.T) {
	var gotOnShutdown = make(chan struct{}, 1)

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "127.0.0.1:0",
		Handler: mux,
	}
	server.RegisterOnShutdown(func() { gotOnShutdown <- struct{}{} })
	go func() {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("signal")
			select {
			case <-sig:
				server.Shutdown(context.Background())
			}
		}
	}()
	go func() {
		time.Sleep(5 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	}()
	server.ListenAndServe()
	select {
	case <-gotOnShutdown:
		fmt.Println("i am close")
	case <-time.After(5 * time.Second):
		t.Errorf("onShutdown callback not called, RegisterOnShutdown broken?")
	}
}
