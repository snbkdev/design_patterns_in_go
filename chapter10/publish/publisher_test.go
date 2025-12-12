package main

type mockSubscriber struct {
	notifyTestingFunc func(msg interface{})
	closeTestingFunc func()
}

func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}

func (m *mockSubscriber) Notify(msg interface{}) error {
	m.notifyTestingFunc(msg)
	return nil
}