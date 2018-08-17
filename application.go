package main

import (
	"fmt"
	"github.com/just1689/fun-with-chan/state"
	"time"
)

func main() {

	fmt.Println("Starting")

	topic := state.NewTopic("WORK")

	go func() {
		for i := 1; i <= 10000; i++ {
			msg := fmt.Sprint(i)
			topic.PutItem(msg)
		}
	}()

	createConsumer(topic, "A")
	createConsumer(topic, "B")
	createConsumer(topic, "C")
	createConsumer(topic, "D")
	createConsumer(topic, "E")

	time.Sleep(10 * time.Second)
}

func createConsumer(topic *state.Topic, ID string) {
	c := topic.Subscribe(ID)
	go func() {
		for item := range c {
			fmt.Println("<-", item.Msg, "says consumer", ID)
			topic.CompletedItem(state.DoneMessage{ConsumerID: ID, ItemID: item.ID})
		}
	}()

}
