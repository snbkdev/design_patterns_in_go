package main

type Subscriber interface {
    Notify(interface{}) error
    Close()
}

type Publisher interface {
    start()
    AddSubscriberCh() chan<- Subscriber
    RemoveSubscriberCh() chan<- Subscriber
    PublishingCh() chan<- interface{}
    Stop()
}

type publisher struct {
    subscribers  []Subscriber
    addSubCh     chan Subscriber
    removeSubCh  chan Subscriber
    in           chan interface{}
    stop         chan struct{}
}

func NewPublisher() Publisher {
    p := &publisher{
        subscribers:  make([]Subscriber, 0),
        addSubCh:     make(chan Subscriber),
        removeSubCh:  make(chan Subscriber),
        in:           make(chan interface{}),
        stop:         make(chan struct{}),
    }
    p.start()
    return p
}

func (p *publisher) start() {
    go func() {
        for {
            select {
            case msg := <-p.in:
                for _, sub := range p.subscribers {
                    _ = sub.Notify(msg)
                }
            case sub := <-p.addSubCh:
                p.subscribers = append(p.subscribers, sub)
            case sub := <-p.removeSubCh:
                for i, candidate := range p.subscribers {
                    if candidate == sub {
                        p.subscribers = append(p.subscribers[:i], 
                            p.subscribers[i+1:]...)
                        break
                    }
                }
            case <-p.stop:
                return
            }
        }
    }()
}

func (p *publisher) AddSubscriberCh() chan<- Subscriber {
    return p.addSubCh
}

func (p *publisher) RemoveSubscriberCh() chan<- Subscriber {
    return p.removeSubCh
}

func (p *publisher) PublishingCh() chan<- interface{} {
    return p.in
}

func (p *publisher) Stop() {
    close(p.stop)
}