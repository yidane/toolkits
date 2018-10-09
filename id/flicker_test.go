package id

import (
	"database/sql"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func newFlicker() *Flicker {
	db, err := sql.Open("mysql", "root:sasasasasa@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	flicker := NewFlicker(db)
	err = flicker.SetTableName("AllSequence").SetAutoMigration(true).Ready()
	if err != nil {
		panic(err)
	}

	return flicker
}

func TestFlicker_NewID(t *testing.T) {
	flicker := newFlicker()
	id, err := flicker.NewID()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(id)
}

func TestFlicker_NewID_Thread(t *testing.T) {
	flicker := newFlicker()
	flicker.SetBatchSize(10)
	wg := sync.WaitGroup{}
	newId := func() int64 {
		id, err := flicker.NewID()
		if err != nil {
			panic(err)
		}

		wg.Done()
		return id
	}

	for i := 0; i < 2000; i++ {
		go newId()
		wg.Add(1)
	}

	wg.Wait()
}

func TestFlicker_NewID_Thread_Check(t *testing.T) {
	flicker := newFlicker()
	flicker.SetBatchSize(10)
	wg := sync.WaitGroup{}
	var ids = make(map[int64]int)
	ch := make(chan int64)
	newId := func() int64 {
		id, err := flicker.NewID()
		if err != nil {
			panic(err)
		}

		wg.Done()

		ch <- id

		return id
	}

	for i := 0; i < 2000; i++ {
		go newId()
		wg.Add(1)
	}

	go func() {
		for id := range ch {
			if _, ok := ids[id]; ok {
				ids[id] += 1
			} else {
				ids[id] = 1
			}
		}

	}()

	wg.Wait()

	//check only one
	for id, tats := range ids {
		if tats > 1 {
			fmt.Println(id, ":", tats)
			t.Fatal("ids are duplicate")
		}
	}

	//check continue
	var idArr []float64
	for id := range ids {
		idArr = append(idArr, float64(id))
	}

	sort.Float64s(idArr)

	for i := 0; i < len(idArr)-1; i++ {
		if idArr[i]+1 != idArr[i+1] {
			fmt.Println(idArr[i]+1, ":", idArr[i+1])
			t.Fatal("ids are not continuous")
		}
	}
}

func BenchmarkFlicker_NewID(b *testing.B) {
	flicker := newFlicker()
	flicker.SetBatchSize(100)
	beginId, _ := flicker.NewID()
	var lastId int64
	for i := 0; i < b.N; i++ {
		id, _ := flicker.NewID()
		atomic.CompareAndSwapInt64(&lastId, lastId, id)
	}

	if beginId+int64(b.N) != int64(b.N) {
		b.Fatal(fmt.Sprintf("%v + %v = %v = %v", beginId, b.N, lastId, beginId+int64(b.N)))
	}
}
