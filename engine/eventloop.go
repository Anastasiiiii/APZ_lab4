package engine

type Handler interface {
	Post(c Command)
}

type EventLoop struct {
	q *CommandQueue

	stopped    bool
	stopSignal chan struct{}
}

func (l *EventLoop) Start() {
	l.q = &CommandQueue{
		notEmptySignal: make(chan struct{}),
	}
	l.stopSignal = make(chan struct{})
	go func() {
		for !l.stopped || !l.q.empty() {
			cmd := l.q.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()
}

func (l *EventLoop) Post(cmd Command) {
	l.q.push(cmd)
}

func (l *EventLoop) AwaitFinish() {
	l.Post(StopCommand{})
	l.stopped = true
	<-l.stopSignal
}
