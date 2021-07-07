package queue

import (
	"container/list"
	"reflect"
	"sync"
)

type BaseQueue struct {
	typ  reflect.Type
	Lock sync.Mutex
	List *list.List
}

func NewBaseQueue(typ reflect.Type) *BaseQueue {
	return &BaseQueue{
		typ:  typ,
		List: list.New(),
	}
}

func (p *BaseQueue) Offer(e interface{}) *list.Element {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	if p.typ != reflect.TypeOf(e){
		return nil
	}

	return p.List.PushBack(e)
}
func (p *BaseQueue) Remove(e *list.Element) interface{}{
	p.Lock.Lock()
	defer p.Lock.Unlock()
	if p.typ != nil && p.typ != reflect.TypeOf(e.Value) {
		return nil
	}

	return p.List.Remove(e)
}
func (p *BaseQueue) Poll() *list.Element{
	if p.List.Len() == 0 {
		return nil
	}

	p.Lock.Lock()
	defer p.Lock.Unlock()

	if p.List.Len() == 0 {
		return nil
	}
	e := p.List.Front()
	p.List.Remove(e)
	return e
}
func (p *BaseQueue) Peek() *list.Element{
	if p.List.Len() == 0 {
		return nil
	}

	p.Lock.Lock()
	defer p.Lock.Unlock()

	if p.List.Len() == 0 {
		return nil
	}
	return p.List.Front()
}

func (p *BaseQueue) Size() int {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	return p.List.Len()
}