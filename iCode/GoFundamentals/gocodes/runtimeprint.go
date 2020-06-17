package gocodes

import (
	"fmt"
	"runtime"
)

func Runtimeprint() {
	fmt.Println("Hello from", runtime.GOOS)
}
