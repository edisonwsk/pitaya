package queue

import "container/list"

type Queue interface {
	Offer(e interface{}) *list.Element
	Remove(e *list.Element) interface{}
	Poll() *list.Element
	Peek() *list.Element
	Size() int
}
