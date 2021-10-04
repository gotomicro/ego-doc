package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrapError(t *testing.T) {
	fmt.Printf("%+v\n", wrapError())
	fmt.Println(wrapError())
	fmt.Println(errors.Unwrap(wrapError()))
	fmt.Println(errors.Is(wrapError(), Err1))
}

const hello = errorstring("hello")

type errorstring string

func (a errorstring) Error() string {
	return string(a)
}

func TestErrorIsStr(t *testing.T) {
	fmt.Println(hello == errorstring("hello"))
}

var hellostructpointer = &errorstruct{"hello", "haha"}
var hellostruct = errorstruct{"hello", "haha"}

type errorstruct struct {
	info string
	code string
}

func (a errorstruct) Error() string {
	return fmt.Sprintf("%v%v", a.info, a.info)
}

func TestErrorIsStruct(t *testing.T) {
	fmt.Println(hellostruct == errorstruct{"hello", "haha"})
	fmt.Println(hellostructpointer == &errorstruct{"hello", "haha"})
	pp := &errorstruct{}
	fmt.Println(errors.As(hellostructpointer, &pp))
	fmt.Println(errors.Is(&errorstruct{"hello", "haha"}, hellostructpointer))
}
