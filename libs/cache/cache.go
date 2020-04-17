package cache

import (
	"container/list"
	"sync"
	"time"
)

type Cache struct {
	rw        *sync.RWMutex
	size      int
	entryList *list.List
	items     *sync.Map
}

type entry struct {
	key     interface{}
	value   interface{}
	expired int64
}

func NewCache(size int) *Cache {
	return &Cache{size: size, entryList: list.New(), items: new(sync.Map)}
}

func (lru *Cache) Get(key interface{}) (val interface{}, ok bool) {
	if item0, ok := lru.items.Load(key); ok {
		item := item0.(*list.Element)
		val = item.Value.(*entry).value
		if item.Value.(*entry).expired < time.Now().Unix() {
			lru.removeElement(item)
			return nil, false
		}
		if val != nil {
			lru.entryList.MoveToFront(item)
			return val, true
		}
		lru.removeElement(item)
	}
	return nil, false
}

func (lru *Cache) Put(key interface{}, value interface{}, expired int64) {
	if item0, ok := lru.items.Load(key); ok {
		item := item0.(*list.Element)
		lru.entryList.MoveToFront(item)
		ent := item.Value.(*entry)
		ent.value = value
		ent.expired = time.Now().Unix() + expired
		return
	}
	element := lru.entryList.PushFront(&entry{key: key, value: value, expired: time.Now().Unix() + expired})
	lru.items.Store(key, element)
	if lru.size > 0 && lru.size < lru.entryList.Len() {
		lru.removeOldItem()
	}
}

func (lru *Cache) Remove(key interface{}) {
	if item0, ok := lru.items.Load(key); ok && item0 != nil {
		item := item0.(*list.Element)
		if item != nil {
			lru.removeElement(item)
		}

	}
}

func (lru *Cache) removeOldItem() {
	lastItem := lru.entryList.Back()
	if lastItem != nil {
		lru.removeElement(lastItem)
	}
}

func (lru *Cache) removeElement(item *list.Element) {
	lru.entryList.Remove(item)
	if item.Value == nil {
		return
	}
	ent := item.Value.(*entry)
	lru.items.Delete(ent.key)

}

// Resize changes the cache size.
func (lru *Cache) Resize(size int) (evicted int) {
	diff := lru.entryList.Len() - size
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff && size > 0; i++ {
		lru.removeOldItem()
	}
	lru.size = size
	return diff
}

var defaultCache *Cache = NewCache(1000)

func Get(key interface{}) (r interface{}, flag bool) {
	return defaultCache.Get(key)
}

func Put(key interface{}, value interface{}, expired int64) {
	defaultCache.Put(key, value, expired)
}
