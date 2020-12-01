package util

import (
	"runtime"
	"time"
)

func GenerateCPULoadFor(d time.Duration) {
	done := make(chan bool)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(d)
	close(done)
}
