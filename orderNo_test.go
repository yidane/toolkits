package toolkits

import (
	"fmt"
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

	fmt.Println(len(nos))
}
