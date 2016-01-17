package gokicadlib

import "fmt"

type TimeStamp int64

func (t TimeStamp) String() string {
	return fmt.Sprintf("%8.8lX", t)
}
