package datastructslice

import "fmt"

type Message struct {
	Content string
	Urgency int
}

type Queue struct{}

func GetListItems() []Message {
	var messages []Message

	messages = []Message{
		{
			Content: "This is Message1",
			Urgency: 10,
		}, {
			Content: "This is Message2",
			Urgency: 20,
		}, {
			Content: "This is Message3",
			Urgency: 30,
		}, {
			Content: "This is Message4",
			Urgency: 40,
		},
	}
	return messages
}

func Pop(msg []Message) *[]Message {
	var message []Message
	var index int
	var list []int

	for _, v := range msg {
		list = append(list, v.Urgency)
	}
	big := list[0]
	for key, value := range list {
		if value > big {
			big = value
			index = key
		}
	}
	message = msg[:index]
	return &message
}

func main() {
	msg := GetListItems()

	result := Pop(msg)
	fmt.Println("Result is", result)
	result = Pop(*result)
	fmt.Println("Result is", result)
	result = Pop(*result)
	fmt.Println("Result is", result)
	result = Pop(*result)
	fmt.Println("Result is", result)

}
