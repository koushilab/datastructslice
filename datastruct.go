package datastructslice

type Message struct {
	Content string
	Urgency int
}

type Queue struct{}

func (q *Queue) Add(m Message) {

}

func (q *Queue) Pop() *Message {
	var message Message
	return &message
}
