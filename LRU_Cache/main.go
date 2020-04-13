package main

//

// Inefficient
// O(N) time and space

import (
	"container/list"
)

// Efficient
// O(1) time and space

type LRUCache struct {
	store     map[int]int
	capacity  int
	order     *list.List
	order_map map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	store := make(map[int]int, capacity)
	order := list.New()
	order_map := make(map[int]*list.Element, capacity)
	return LRUCache{store: store, capacity: capacity, order: order, order_map: order_map}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.order_map[key]; ok {
		this.order.MoveToFront(elem)
		return this.store[key]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if len(this.order_map) == this.capacity {
		if _, ok := this.store[key]; !ok {
			temp := this.order.Remove(this.order.Back())
			delete(this.store, temp.(int))
			delete(this.order_map, temp.(int))
		}

	}
	this.store[key] = value
	if val, ok := this.order_map[key]; ok {
		this.order.MoveToFront(val)
	} else {
		this.order_map[key] = this.order.PushFront(key)
	}

}
