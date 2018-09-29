package id

import (
	"fmt"
	"sync"
	"time"
)

const (
	epoch              int64 = 1524199200000 //2018-04-20 12:40:000. The moment when my son born!
	workerIDBits       uint  = 5
	dataCenterIDBits   uint  = 5
	maxWorkerID        uint  = -1 ^ (-1 << workerIDBits)
	maxDataCenterID    uint  = -1 ^ (-1 << dataCenterIDBits)
	sequenceBits       uint  = 12
	workerIdShift            = sequenceBits
	dataCenterIdShift        = sequenceBits + workerIDBits
	timestampLeftShift       = sequenceBits + workerIDBits + dataCenterIDBits
	sequenceMask       uint  = -1 ^ (-1 << sequenceBits)
)

type Snowflake struct {
	dataCenterID  uint
	workerID      uint
	sequence      uint
	lastTimestamp int64
	lock          sync.Mutex
}

func NewSnowflake(dataCenterID, workerID, sequence uint) (*Snowflake, error) {
	if workerID > maxWorkerID || workerID < 0 {
		return nil, fmt.Errorf("worker Id can't be greater than %v or less than 0", maxWorkerID)
	}

	if dataCenterID > maxDataCenterID || dataCenterID < 0 {
		return nil, fmt.Errorf("datacenter Id can't be greater than %v or less than 0", maxDataCenterID)
	}

	rtnSnowflake := Snowflake{
		dataCenterID: dataCenterID,
		workerID:     workerID,
		sequence:     sequence,
		lock:         sync.Mutex{},
	}

	return &rtnSnowflake, nil
}

func (snowflake *Snowflake) GetDataCenterID() uint {
	return snowflake.dataCenterID
}

func (snowflake *Snowflake) GetWorkerID() uint {
	return snowflake.workerID
}

func getUnixNano() int64 {
	return time.Now().UnixNano()
}

//NewID generate a new id
func (snowflake *Snowflake) NewID() (int64, error) {
	snowflake.lock.Lock()
	defer snowflake.lock.Unlock()

	currentTime := getUnixNano()
	var sequence uint

	switch {
	case currentTime < snowflake.lastTimestamp:
		return 0, fmt.Errorf("Clock moved backwards.  Refusing to generate id for %v milliseconds", snowflake.lastTimestamp-currentTime)
	case currentTime == snowflake.lastTimestamp:
		sequence = (snowflake.sequence + 1) & sequenceMask
		if sequence == 0 {
			currentTime = getUnixNano()
			for currentTime <= snowflake.lastTimestamp {
				currentTime = getUnixNano()
			}
		}
	case currentTime > snowflake.lastTimestamp:
		sequence = 0
	}

	snowflake.lastTimestamp = currentTime
	return (currentTime-epoch)<<timestampLeftShift |
		(int64(snowflake.dataCenterID) << dataCenterIdShift) |
		(int64(snowflake.workerID) << workerIdShift) |
		int64(sequence), nil
}
