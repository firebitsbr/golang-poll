package main

import (
	"golang-poll/api"
	"runtime"
)

func main() {
	// Max cpu
	runtime.GOMAXPROCS(runtime.NumCPU())

	// bootstrap
	api.Bootstrap()
}