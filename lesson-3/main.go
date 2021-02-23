package main

import (
	"GB/lesson-3/eventchannel"
	"GB/lesson-3/prototype"
	"GB/lesson-3/singleton"
	"fmt"
)

func main() {
	s1 := singleton.NewSingleStruct()
	s2 := singleton.NewSingleStruct()
	fmt.Println(s1 == s2)

	p1 := prototype.NewPrototypeStruct()
	p2 := p1.Clone()
	fmt.Println(p1.Compare(p2))

	eventHandler := eventchannel.NewEventHandler()
	pub := eventchannel.NewPublisher(eventHandler, []string{"key1", "key2"})
	sub1 := eventchannel.NewSubscriber(eventHandler, "key1")
	sub2 := eventchannel.NewSubscriber(eventHandler, "key2")
	go func() {
		pub.Publish([]byte("Hello GO!"))
	}()

	fmt.Println(string(sub1.GetMessage()))
	fmt.Println(string(sub2.GetMessage()))
}
