package network

import (
	"sync"
)

const (
	defaultPopListSize = 32
)

//PackCache cls cacheItem
type PackCache struct {
	Pack    PackInf
	Session *TCPSession
}

//PackQueue cls queue
type PackQueue struct {
	sync.Mutex
	items    []*PackCache
	readPos  int
	writePos int
	nLen     int
}

//NewPackQueue interface of network
func NewPackQueue(capacity int) *PackQueue {
	q := &PackQueue{items: make([]*PackCache, 0, capacity)}
	return q
}

//Push put item to queue
func (q *PackQueue) Push(item *PackCache) error {
	q.Lock()
	defer q.Unlock()
	q.items = append(q.items, item)
	return nil
}

//Pop pop item from queue
func (q *PackQueue) Pop() *PackCache {
	q.Lock()
	defer q.Unlock()
	if len(q.items) > 0 {
		item := q.items[0]
		q.items = q.items[1:]
		return item
	}
	return nil
}

//PopList pop serval items
func (q *PackQueue) PopList(size int) []*PackCache {
	q.Lock()
	defer q.Unlock()
	if size <= 0 {
		size = defaultPopListSize
	}
	wLen := len(q.items)
	if wLen >= size {
		items := q.items[0:size]
		q.items = q.items[size:]
		return items
	}

	items := q.items[0:wLen]
	q.items = q.items[wLen:]
	return items

}

//PopAll pop all items
func (q *PackQueue) PopAll() []*PackCache {
	q.Lock()
	defer q.Unlock()
	if len(q.items) == 0 {
		return nil
	}
	items := q.items[0:]
	q.items = q.items[len(q.items):]
	return items
}

//Dispose dispose queue
func (q *PackQueue) Dispose() {
	q.Lock()
	defer q.Unlock()
	q.items = make([]*PackCache, 0, 0)
}
