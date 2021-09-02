package main

import (
	"context"
	"testing"

	"github.com/gotomicro/ego/examples/helloworld"
	"github.com/stretchr/testify/assert"
)

func TestGreeter_SayHello(t *testing.T) {
	s := Greeter{}
	info, err := s.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "askuy",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Hello EGO! I am askuy", info.Message)
}
