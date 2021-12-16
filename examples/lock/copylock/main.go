package main

import (
	"fmt"
	"sync"
)

type User struct {
	sync.Mutex

	name string
}

func main() {
	u1 := &User{name: "test"}
	u1.Lock()
	defer u1.Unlock()

	tmp := *u1
	u2 := &tmp
	// u2.Mutex = sync.Mutex{} // 没有这一行就会死锁

	fmt.Printf("%#p\n", u1)
	fmt.Printf("%#p\n", u2)

	u2.Lock()
	defer u2.Unlock()
}
