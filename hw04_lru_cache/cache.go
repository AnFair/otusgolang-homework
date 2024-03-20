package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type Pair struct {
	key   Key
	value interface{}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	item, ok := cache.items[key]
	if ok {
		cache.queue.MoveToFront(item)
		item.Value = Pair{key, value}
	} else {
		if cache.capacity == cache.queue.Len() {
			last := cache.queue.Back()
			cache.queue.Remove(last)
			delete(cache.items, last.Value.(Pair).key)
		}
		lruItem := cache.queue.PushFront(Pair{key, value})
		cache.items[key] = lruItem
	}
	return ok
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := cache.items[key]
	if ok {
		cache.queue.MoveToFront(item)
		return item.Value.(Pair).value, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
