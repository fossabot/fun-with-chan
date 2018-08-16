package state

import (
	"container/ring"
)

type Topic struct {
	Name              string
	Head              *ring.Ring
	Count             int
	CountID           int64
	Incoming          chan string
	Completed         chan int64
	Consumer          *Consumer
	incomingConsumers chan Consumer
}

func NewTopic(name string) *Topic {
	t := Topic{Name: name, Count: 0, CountID: 0}
	t.Incoming = make(chan string)
	t.Completed = make(chan int64)
	t.manageIO()
	return &t
}

func (t *Topic) manageIO() {
	go func() {
		for {
			select {
			case c := <-t.incomingConsumers:
				t.Consumer = &c
				break
			case ID := <-t.Completed:
				t.markDone(ID)
				break
			case in := <-t.Incoming:
				t.handleIn(in)
				break
			}
		}
	}()
}

func (t *Topic) PutItem(msg string) {
	t.Incoming <- msg
}

func (t *Topic) CompletedItem(ID int64) {
	t.Completed <- ID
}

func (t *Topic) Subscribe() chan *Item {
	consumer := Consumer{Idle: true}
	consumer.Channel = make(chan *Item)
	t.incomingConsumers <- consumer
	return consumer.Channel
}

func (t *Topic) handleIn(msg string) {

	t.Count++

	if t.Count == 1 {
		r := ring.New(1)
		t.Head = r
		t.Head.Value = &Item{ID: t.CountID, Msg: msg, Busy: false}
		return
	}

	r := ring.New(1)
	r.Value = &Item{ID: t.CountID, Msg: msg, Busy: false}
	r.Link(t.Head)

	t.work()

}

func (t *Topic) canWork() bool {

	if t.Consumer == nil || t.Consumer.Idle == false {
		return false
	}

	if t.Count == 0 {
		return false
	}
	if t.Consumer == nil {
		return false
	}

	item := t.Head.Value.(*Item)
	ca := item.Busy == false
	return ca

}

func (t *Topic) work() {
	if !t.canWork() {
		return
	}

	item := t.Head.Value.(*Item)
	t.Consumer.Channel <- item
	item.Busy = true
	t.Consumer.Idle = false

}
func (t *Topic) markDone(ID int64) {
	r := find(t.Head, ID)

	n := t.Head.Next()

	removed := r.Prev().Unlink(1)

	if t.Head == removed {
		t.Head = n
	}

	t.Consumer.Idle = true
	t.Count--

	t.work()
}
