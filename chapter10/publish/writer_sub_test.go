package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

func TestStdoutPrinter(t *testing.T) {}

// func TestWriter(t *testing.T) {
// 	sub := NewWriterSubscriber(0, nil)
// }

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	if out == nil {
		out = os.Stdout
	}

	s := &writerSubscriber{
		id: id,
		in: make(chan interface{}),
		Writer: out,
	}

	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()

	return s
}

type mockWriter struct {
	testingFunc func(string)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}

func TestPublisher(t *testing.T) {
	msg := "Hello"

	p := NewPublisher()

	var wg sync.WaitGroup
	
	sub := &mockSubscriber{
		notifyTestingFunc: func(msg interface{}) {
			defer wg.Done()

			s, ok := msg.(string)
			if !ok {
				t.Fatal(errors.New("Could not assert result!"))
			}

			if s != msg {
				t.Fail()
			}
		},
		closeTestingFunc: func ()  {
			wg.Done()
		},
	}

	p.AddSubscriberCh() <- sub
	wg.Add(1)

	p.PublishingCh() <- msg
	wg.Wait()

	pubCon := p.(*publisher)
	if len(pubCon.subscribers) != 1 {
		t.Error("Unexpected number of subscribers")
	}

	wg.Add(1)
	p.RemoveSubscriberCh() <- sub
	wg.Wait()

	if len(pubCon.subscribers) != 0 {
		t.Error("Expected no subscribers")
	}

	p.Stop()
}