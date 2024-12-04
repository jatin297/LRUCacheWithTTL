package main

import (
	"log"
	"time"
)

// isExpired returns true if time_to_live of node is expired
func (lru *LRU[K, V]) isExpired(node *Node[K, V]) bool {
	if node == nil || node.TTL < 0 {
		return false
	}

	return node.TimeStamp.Add(node.TTL).Before(time.Now())
}

// RunActiveExpirationConcurrently concurrently expire keys based on COUNT_NODES_FOR_EXPIRY_FREQUENCY times per second
func (lru *LRU[K, V]) RunActiveExpirationConcurrently(done chan bool) {
	log.Println("active expiration of keys has been started")

	ticker := time.NewTicker(time.Second / COUNT_NODES_FOR_EXPIRY_FREQUENCY)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			log.Println("stopping keys expiration")
			lru.Wg.Wait()
			return
		case <-ticker.C:
			lru.Wg.Add(1)
			// creating goroutines for expiring keys
			go lru.expireKeys()
		}
	}
}

// RandomKeysWithLiveTTL get all active keys (ttl>=0)
func (lru *LRU[K, V]) RandomKeysWithLiveTTL() []K {

	counter := lru.CountOfNodesForExpiry
	if counter > len(lru.Bucket) {
		counter = len(lru.Bucket)
	}

	keys := []K{}
	for key, node := range lru.Bucket {
		if node.TTL < 0 {
			continue
		}

		keys = append(keys, key)

		counter--
		if counter == 0 {
			break
		}
	}

	return keys
}

func (lru *LRU[K, V]) expireKeys() {
	defer lru.Wg.Done()

	keys := lru.RandomKeysWithLiveTTL()
	n := len(keys)

	if n == 0 {
		return
	}

	expiredKeysCount := 0
	for _, key := range keys {
		node := lru.Bucket[key]
		if lru.isExpired(node) {
			log.Printf("key: %v is expired", key)

			lru.Mutex.Lock()

			delete(lru.Bucket, key)
			lru.Dll.Remove(node)
			lru.Mutex.Unlock()

			expiredKeysCount++
		}
	}

	// if more than 25% of nodes are expired, increase size of expiry nodes capacity
	expiredKeysCountPercentage := expiredKeysCount * 100 / n

	if expiredKeysCountPercentage > EXPIRED_PERCENTAGE_LIMIT {
		lru.Capacity = lru.Capacity * 2
		if lru.Capacity > MAX_COUNT_NODES_FOR_EXPIRY {
			lru.Capacity = MAX_COUNT_NODES_FOR_EXPIRY
		}
	}
}
