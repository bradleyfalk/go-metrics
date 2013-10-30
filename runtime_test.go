package metrics

import (
	"runtime"
	"testing"
	"time"
)

func BenchmarkRuntimeMemStats(b *testing.B) {
	r := NewRegistry()
	RegisterRuntimeMemStats(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CaptureRuntimeMemStatsOnce(r)
	}
}

func TestRuntimeMemStatsBlocking(t *testing.T) {
	ch := make(chan int)
	go testRuntimeMemStatsBlocking(ch)
	//runtime.Gosched()
	var memStats runtime.MemStats
	t0 := time.Now()
	runtime.ReadMemStats(&memStats)
	t1 := time.Now()
	t.Log("i++ during runtime.ReadMemStats:", <-ch)
	go testRuntimeMemStatsBlocking(ch)
	time.Sleep(t1.Sub(t0))
	t.Log("i++ during time.Sleep:", <-ch)
}

func testRuntimeMemStatsBlocking(ch chan int) {
	i := 0
	for {
		select {
		case ch <- i:
			return
		default:
			i++
		}
	}
}