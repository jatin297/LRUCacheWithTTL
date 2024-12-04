package main

import "time"

// if node already exists, we check if node expired or not
// if expired remove it, create new
// if not expired, update it and timestamp to current
// if key not exists, check if free space in dll
// if space present insert at front, else remove last

func (lru *LRU[K, V]) Set(key K, value V, ttl time.Duration) {
	lru.Mutex.Lock()
	defer lru.Mutex.Unlock()

	node, exists := lru.Bucket[key]
	// if key exists, update value and move node to front
	if exists {
		if lru.isExpired(node) {
			delete(lru.Bucket, node.Key)
			lru.Dll.Remove(node)

			// add new key-value node in lru bucket
			newNode := NewNode(key, value, ttl)
			lru.Bucket[key] = newNode

			// move newly added node to front
			lru.Dll.Prepend(newNode)
			return
		}

		node.Value = value
		node.TimeStamp = time.Now()
		lru.Dll.Remove(node)
		lru.Dll.Prepend(node)
	} else {

		// check if bucket capacity exceeded, remove least recently node (last)
		if len(lru.Bucket) >= lru.Capacity {
			delete(lru.Bucket, lru.Dll.Tail.Key)
			lru.Dll.Remove(lru.Dll.Tail)
		}

		newNode := NewNode(key, value, ttl)
		lru.Bucket[key] = newNode

		// add new node at front of dll
		lru.Dll.Prepend(newNode)
	}
}
