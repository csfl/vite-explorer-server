package util

import (
	"github.com/vitelabs/go-vite/common/types"
	"time"
	"sync"
)

var timeLimit = 3 // timestamp是秒级
var apiAccessMap = make(map[types.Address]int64)
var checkLock sync.Mutex

func CheckApiLimit (addr types.Address) bool {
	checkLock.Lock()
	defer  checkLock.Unlock()

	currentTimestamp := time.Now().Unix()
	if lastTimestamp, ok := apiAccessMap[addr]; ok && currentTimestamp - lastTimestamp <= 3 {
		return false
	}
	apiAccessMap[addr] = currentTimestamp
	return true
}
