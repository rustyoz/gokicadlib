package gokicadlib

import (
	"fmt"

	"time"
)

type TimeStamp int64

func (t TimeStamp) String() string {
	s := fmt.Sprintf("%X", int64(t))
	return s
}

func (t *TimeStamp) Stamp() *TimeStamp {
	var nt TimeStamp
	nt = TimeStamp(time.Now().UnixNano())
	t = &nt
	return t
}
