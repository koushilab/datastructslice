package datastructslice

import (
	"testing"
)


func TestPop(t *testing.T) {
	var messages []Message
	message1 := Message{
		Content: "This is Message1",
		Urgency: 10,
	}
	message2 := Message{
		Content: "This is Message2",
		Urgency: 20,
	}
	messages = append(messages, message1)
	messages = append(messages, message2)

	result := Pop(messages)

	expected := append(messages, message1)

	for _, v := range *result {
		if v != message1 {
			t.Errorf("Test failed Expected : %v, Got: %v", expected, v)
		}
	}

}



func TestPop1(t *testing.T) {

	getItems := GetListItems()

	result := Pop(getItems)

	expected := &[]Message{
		{
			Content: "This is Message1",
			Urgency: 10,
		}, {
			Content: "This is Message2",
			Urgency: 20,
		}, {
			Content: "This is Message3",
			Urgency: 30,
		},
	}

	if result == expected {
		t.Errorf("Test failed Expected : %v, Got: %v", expected, result)
	}
}
