package main

import "time"

// Get: returns node from dll and move it to front
func (lru *LRU[K, V]) Get(key K) (V, bool) {

	if node, ok := lru.Bucket[key]; ok {

		if lru.isExpired(node) {
			lru.Mutex.Lock()

			delete(lru.Bucket, node.Key)
			lru.Dll.Remove(node)

			lru.Mutex.Unlock()

			return *new(V), false
		}

		node.TimeStamp = time.Now()
		lru.Mutex.Lock()

		lru.Dll.Remove(node)
		lru.Dll.Prepend(node)

		lru.Mutex.Unlock()
		return node.Value, true
	}

	return *new(V), false
}
