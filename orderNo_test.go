package toolkits

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestCreateOrderNo(t *testing.T) {
	CreateOrderNo(nil)
	CreateOrderNo(func() *time.Time {
		now := time.Now()
		return &now
	})
}

func TestCreateOrderNoBench(t *testing.T) {
	nos := make(map[string]int)
	lock := sync.RWMutex{}
	appendNo := func(no string) {
		lock.Lock()
		defer lock.Unlock()

		if t, f := nos[no]; f {
			nos[no] = t + 1
		} else {
			nos[no] = 1
		}
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}
	total := 20000000
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			no := CreateOrderNo(nil)
			appendNo(no)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(len(nos))
}

func BenchmarkCreateOrderNo(b *testing.B) {
	nos := make(map[string]int)
	lock := sync.RWMutex{}
	appendNo := func(no string) {
		lock.Lock()
		defer lock.Unlock()

		if t, f := nos[no]; f {
			nos[no] = t + 1
		}
	}

	for i := 0; i < b.N; i++ {
		no := CreateOrderNo(nil)
		appendNo(no)
	}
}
