package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrapError(t *testing.T) {
	fmt.Println(wrapError())
	fmt.Println(errors.Unwrap(wrapError()))
	fmt.Println(errors.Is(wrapError(), Err1))
}
