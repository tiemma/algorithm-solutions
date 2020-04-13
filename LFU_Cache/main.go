package main

// https://leetcode.com/problems/lfu-cache/

import (
	"container/list"
	"math"
)

type LFUCache struct {
	store    map[int]int
	capacity int

	order     *list.List
	order_map map[int]*list.Element

	count     map[int][2]int // 0 - key count, 1 - overall count for the cache
	curr_idx  int
	max_count int
}

func Constructor(capacity int) LFUCache {
	if capacity <= 0 {
		return LFUCache{}
	}
	store := make(map[int]int, capacity)
	count := make(map[int][2]int, capacity)
	count[0] = [2]int{math.MaxInt64, math.MaxInt64}

	order := list.New()
	order_map := make(map[int]*list.Element, capacity)

	return LFUCache{store: store, capacity: capacity, order: order, order_map: order_map, count: count}
}

func (this *LFUCache) SetCURRIdx() {
	for k, val := range this.count {
		if val[0] < this.count[this.curr_idx][0] {
			this.curr_idx = k
		} else if val[0] == this.count[this.curr_idx][0] {
			if val[1] < this.count[this.curr_idx][1] {
				this.curr_idx = k
			}
		}
	}
}

func (this *LFUCache) SetLRUEntry(key int) {
	this.max_count += 1
	this.count[key] = [2]int{this.count[key][0] + 1, this.max_count}
}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}
	if elem, ok := this.order_map[key]; ok {
		this.SetLRUEntry(key)
		this.order.MoveToBack(elem)
		return this.store[key]
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}
	if len(this.order_map) == this.capacity {
		if _, ok := this.store[key]; !ok {
			this.SetCURRIdx()
			temp := this.order.Remove(this.order_map[this.curr_idx])
			delete(this.store, temp.(int))
			delete(this.order_map, temp.(int))
			delete(this.count, temp.(int))
			this.curr_idx = 0
		}
	}

	if val, ok := this.order_map[key]; ok {
		this.order.MoveToBack(val)
	} else {
		this.order_map[key] = this.order.PushFront(key)
		this.count[key] = [2]int{-1, this.max_count}
	}
	this.SetLRUEntry(key)
	this.store[key] = value
}
