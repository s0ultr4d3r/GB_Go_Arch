/**
TODO: Mutex for flow-safe
**/

package lru

import (
	"GB/lesson-5/doublelinkedlist"
)

type Item struct {
	Key   string
	Value interface{}
}

type LRU struct {
	capacity int
	items    map[string]*doublelinkedlist.Elem
	queue    *doublelinkedlist.List
}

func NewLru(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*doublelinkedlist.Elem),
		queue:    doublelinkedlist.New(),
	}
}

func (c *LRU) Set(key string, value interface{}) bool {
	if element, exists := c.items[key]; exists == true {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return true
	}

	if c.queue.Len() == c.capacity {
		c.purge()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.Key] = element

	return true
}

func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
	}
}

func (c *LRU) Get(key string) interface{} {
	element, exists := c.items[key]
	if exists == false {
		return nil
	}
	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value
}
