package main

import (
	"container/list"
	"fmt"
)

type node struct {
	key int
	val int
}

type LRUCache struct {
	mapping map[int]*list.Element
	cap     int
	dlist   *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		mapping: make(map[int]*list.Element),
		cap:     capacity,
		dlist:   list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.mapping[key]; ok {
		this.dlist.MoveToFront(ele)
		return ele.Value.(node).val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.mapping[key]; ok {
		ele.Value = node{key: key, val: value}
		this.dlist.MoveToFront(ele)
	} else {
		newEle := node{key: key, val: value}
		this.dlist.PushFront(newEle)
		this.mapping[key] = this.dlist.Front()

		if len(this.mapping) > this.cap {
			last := this.dlist.Back()
			this.dlist.Remove(last)
			delete(this.mapping, last.Value.(node).key)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(5)
	obj.Put(2, 2)
	fmt.Println(obj)

}
