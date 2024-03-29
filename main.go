package main

import (
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)

	log.Infof("Hello World from %s/%s", runtime.GOOS, runtime.GOARCH)
}

func add(a, b int) int {
	return a + b
}
