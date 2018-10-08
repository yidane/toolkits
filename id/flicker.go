package id

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

const defaultBatchSize int64 = 10000
const defaultTableName string = "flicker_sequence"

type Flicker struct {
	stub      string
	batchSize int64
	sequence  int64
	maxId     int64
	currentId int64
	ready     bool
	tableName string
	lock      sync.RWMutex
	db        *sql.DB
}

//NewFlicker create instance of Flicker
func NewFlicker(db *sql.DB) *Flicker {
	flicker := &Flicker{
		stub:      defaultStub(),
		batchSize: defaultBatchSize,
		db:        db,
		lock:      sync.RWMutex{},
	}

	return flicker
}

func (flicker *Flicker) SetStub(stub string) *Flicker {
	if strings.TrimSpace(stub) == "" {
		stub = defaultStub()
	}

	flicker.stub = stub

	return flicker
}

func (flicker *Flicker) SetBatchSize(batchSize int64) *Flicker {
	if batchSize <= 0 {
		batchSize = defaultBatchSize
	}

	flicker.batchSize = batchSize

	return flicker
}

func (flicker *Flicker) SetTableName(tableName string) *Flicker {
	if strings.TrimSpace(tableName) == "" {
		tableName = defaultTableName
	}

	flicker.tableName = tableName

	return flicker
}

//Ready check all config is ready
func (flicker *Flicker) Ready() error {
	if flicker.ready {
		return nil
	}

	if flicker.db == nil {
		return fmt.Errorf("mysql connection is nil")
	}

	var dri interface{} = flicker.db.Driver()
	if _, ok := dri.(*mysql.MySQLDriver); !ok {
		return fmt.Errorf("database type must be mysql")
	}

	if err := flicker.db.Ping(); err != nil {
		return err
	}

	flicker.ready = true

	return nil
}

func (flicker *Flicker) GetStub() string {
	return flicker.stub
}

func (flicker *Flicker) GetBatchSize() int64 {
	return flicker.batchSize
}

func (flicker *Flicker) GetSequence() int64 {
	return flicker.sequence
}

func (flicker *Flicker) GetMaxID() int64 {
	return flicker.maxId
}

func (flicker *Flicker) GetDB() *sql.DB {
	return flicker.db
}

func defaultStub() string {
	hostName, err := os.Hostname()

	if err != nil {
		log.Println(err)
	}

	return hostName
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
	if !flicker.ready {
		return 0, fmt.Errorf("the instance of Flicker isn't ready,please call the func Ready and check the result first")
	}

	flicker.lock.Lock()
	defer flicker.lock.Unlock()

	atomic.AddInt64(&flicker.currentId, 1)

	return flicker.currentId, nil
}
