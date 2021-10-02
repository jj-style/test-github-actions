package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("I am running on", runtime.GOOS, runtime.GOARCH)
}
