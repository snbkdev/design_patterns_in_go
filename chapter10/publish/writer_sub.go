package main

import (
	"fmt"
	"io"
	"time"
)

type writerSubscriber struct {
	in chan interface{}
	id int
	Writer io.Writer
}

func (s *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()

	select {
	case s.in <- msg:
	case <- time.After(time.Second):
		err = fmt.Errorf("Timeout\n")
	}

	return
}

func (s *writerSubscriber) Close() {
	close(s.in)
}