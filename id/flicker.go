package id

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"sync/atomic"
)

type Flicker struct {
	stub      string
	batchSize int64
	sequence  int64
	maxId     int64
	currentId int64
	lock      sync.RWMutex
	db        sql.DB
}

//NewFlicker create instance of Flicker
func NewFlicker(stub string, db sql.DB, batchSize int64) *Flicker {
	return &Flicker{
		stub:      stub,
		batchSize: batchSize,
		sequence:  0,
		maxId:     0,
		currentId: 0,
		db:        db,
		lock:      sync.RWMutex{},
	}
}

func (flicker *Flicker) newSequence() error {
	result, err := flicker.db.Exec("")
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

//NewID create new id
func (flicker *Flicker) NewID() (int64, error) {
	flicker.lock.Lock()
	defer flicker.lock.Lock()

	atomic.AddInt64(&flicker.currentId, 1)

	return 0, nil
}
