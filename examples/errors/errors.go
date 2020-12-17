package errors

import (
	"errors"
	"fmt"
	"os"
)

var some *os.PathError
var Err1 = errors.New("i am error1")

func getError() error {
	return Err1
}

func wrapError() error {
	return fmt.Errorf("wrap error %w", getError())
}
