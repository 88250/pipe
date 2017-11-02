package util

import (
	"time"
)

func CurrentMillisecond() uint {
	return uint(time.Now().UnixNano() / int64(time.Millisecond))
}
