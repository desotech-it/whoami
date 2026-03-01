package util

import (
	"context"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	interval = 50 * time.Millisecond
)

var (
	cpuCancel context.CancelFunc
	memCancel context.CancelFunc
	cpuMu     sync.Mutex
	memMu     sync.Mutex
)

// GenerateCPULoadFor stresses all CPU cores for the given duration.
// It can be stopped early by calling StopCPUStress.
func GenerateCPULoadFor(d time.Duration) {
	cpuMu.Lock()
	// Cancel any existing stress
	if cpuCancel != nil {
		cpuCancel()
	}
	ctx, cancel := context.WithTimeout(context.Background(), d)
	cpuCancel = cancel
	cpuMu.Unlock()

	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
			}
		}()
	}

	wg.Wait()

	cpuMu.Lock()
	cpuCancel = nil
	cpuMu.Unlock()
}

// StopCPUStress cancels an ongoing CPU stress test.
// Returns true if a stress was running and was stopped.
func StopCPUStress() bool {
	cpuMu.Lock()
	defer cpuMu.Unlock()
	if cpuCancel != nil {
		cpuCancel()
		cpuCancel = nil
		return true
	}
	return false
}

// GenerateHighMemoryUsageFor allocates memory for the given duration.
// It can be stopped early by calling StopMemStress.
func GenerateHighMemoryUsageFor(d time.Duration) {
	memMu.Lock()
	if memCancel != nil {
		memCancel()
	}
	ctx, cancel := context.WithTimeout(context.Background(), d)
	memCancel = cancel
	memMu.Unlock()

	rand.Seed(time.Now().UTC().UnixNano())
	memblock := make([][]byte, 10)
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			memMu.Lock()
			memCancel = nil
			memMu.Unlock()
			return
		default:
			memblock = append(memblock, make([]byte, 1024*1024*16))
			rand.Read(memblock[i])
			time.Sleep(interval)
		}
	}
}

// StopMemStress cancels an ongoing memory stress test.
// Returns true if a stress was running and was stopped.
func StopMemStress() bool {
	memMu.Lock()
	defer memMu.Unlock()
	if memCancel != nil {
		memCancel()
		memCancel = nil
		return true
	}
	return false
}
