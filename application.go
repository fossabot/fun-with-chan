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

		for i := 1; i <= 20; i++ {
			msg := fmt.Sprint(i)
			fmt.Println("Writing: ", msg)
			topic.PutItem(msg)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	createConsumer(topic, "100")
	//createConsumer(topic, "200")

	time.Sleep(10 * time.Second)
}

func createConsumer(topic *state.Topic, ID string) {
	c := topic.Subscribe(ID)
	go func() {
		for item := range c {
			fmt.Println("Message from ", ID, " says: ", item.Msg)
			topic.CompletedItem(state.DoneMessage{ConsumerID: ID, ItemID: item.ID})

		}
	}()

}
