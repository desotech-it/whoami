package util

import (
	"math/rand"
	"runtime"
	"time"
)

const (
	interval = 50 * time.Millisecond
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

func GenerateHighMemoryUsageFor(d time.Duration) {
	done := make(chan bool)

	go func() {
		rand.Seed(time.Now().UTC().UnixNano())
		memblock := make([][]byte, 10)
		for i := 0; true; i++ {
			select {
			case <-done:
				return
			default:
				memblock = append(memblock, make([]byte, 1024*1024*16))
				rand.Read(memblock[i])
				time.Sleep(interval)
			}
		}
	}()

	time.Sleep(d)
	close(done)
}
