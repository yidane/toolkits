package id

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const defaultBatchSize int64 = 10000
const defaultTableName string = "flicker_sequence"

type Flicker struct {
	stub          string
	batchSize     int64
	sequence      int64
	maxId         int64
	currentId     int64
	ready         bool
	autoMigration bool
	tableName     string
	lock          sync.RWMutex
	db            *sql.DB
}

//NewFlicker create instance of Flicker
func NewFlicker(db *sql.DB) *Flicker {
	flicker := &Flicker{
		stub:      defaultStub(),
		batchSize: defaultBatchSize,
		tableName: defaultTableName,
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

//SetAutoMigration set field autoMigration for creating table automatically
func (flicker *Flicker) SetAutoMigration(f bool) *Flicker {
	flicker.autoMigration = f
	return flicker
}

//Ready check all config is ready
func (flicker *Flicker) Ready() error {
	flicker.lock.Lock()
	defer flicker.lock.Unlock()

	if flicker.ready {
		return nil
	}

	if err := flicker.dbPing(); err != nil {
		return err
	}

	if err := flicker.migration(); err != nil {
		return err
	}

	flicker.ready = true

	return nil
}

func (flicker *Flicker) dbPing() error {
	if flicker.db == nil {
		return fmt.Errorf("mysql connection is nil")
	}

	var dri interface{} = flicker.db.Driver()
	if _, ok := dri.(*mysql.MySQLDriver); !ok {
		return fmt.Errorf("database must be mysql")
	}

	return flicker.db.Ping()
}

func (flicker *Flicker) migration() error {
	if flicker.autoMigration {
		createSql := fmt.Sprintf("create table if not exists %s(id bigint auto_increment primary key,stub varchar(100) unique key);", flicker.tableName)
		_, err := flicker.db.Exec(createSql)
		return err
	}

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

func (flicker *Flicker) newSequence() (int64, error) {
	result, err := flicker.db.Exec(fmt.Sprintf("REPLACE INTO %s (stub) VALUES ('%s');", flicker.tableName, flicker.stub))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

//NewID create new id
func (flicker *Flicker) NewID() (int64, error) {
	if !flicker.ready {
		return 0, fmt.Errorf("the instance of Flicker isn't ready,please call the func Ready and check the result first")
	}

	if flicker.currentId >= flicker.maxId {
		flicker.lock.Lock()
		defer flicker.lock.Unlock()

		if flicker.currentId >= flicker.maxId {
			id, err := flicker.newSequence()
			for err != nil {
				time.Sleep(time.Second * 1)
				id, err = flicker.newSequence()
			}

			atomic.SwapInt64(&flicker.sequence, id)
			atomic.SwapInt64(&flicker.currentId, id*flicker.batchSize)
			atomic.SwapInt64(&flicker.maxId, flicker.currentId)
			atomic.AddInt64(&flicker.maxId, flicker.batchSize)
		}
	}
	atomic.AddInt64(&flicker.currentId, 1)

	return flicker.currentId, nil
}
