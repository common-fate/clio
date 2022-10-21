package main

import (
	"sync"

	"github.com/common-fate/clio"
)

func main() {
	clio.Info("hello world")
	clio.Error("test error")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		clio.Info("from a goroutine")
		wg.Done()
	}()
	wg.Wait()
}
