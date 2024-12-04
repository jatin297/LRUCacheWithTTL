package main

import "sync"

const (
	MIN_COUNT_NODES_FOR_EXPIRY = 10
	MAX_COUNT_NODES_FOR_EXPIRY = 100

	// number of times check for expiration should be performed.
	COUNT_NODES_FOR_EXPIRY_FREQUENCY = 10

	// threshold that determines if count of nodes for expiry limit need to be updated
	EXPIRED_PERCENTAGE_LIMIT = 25
)

// LRU: data structures to represent lru cache
type LRU[K comparable, V any] struct {

	// size limit
	Capacity int

	// hashmap map keys with their node in dll
	Bucket map[K]*Node[K, V]

	// doubly linked list to store key value pairs
	// key-values are stored based on their frequency of use
	// order of data in dll: most_recently_used ---- least_recently_used
	Dll *DoublyLinkedList[K, V]

	// CountOfNodesForExpiry: count of nodes with ttl that needs to be checked for expiration
	CountOfNodesForExpiry int

	// mutual exclusion lock to prevent multiple goroutines to access same resource
	Mutex sync.Mutex

	// manage goroutines
	Wg sync.WaitGroup
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		Capacity:              capacity,
		Bucket:                make(map[K]*Node[K, V], capacity),
		Wg:                    sync.WaitGroup{},
		Mutex:                 sync.Mutex{},
		CountOfNodesForExpiry: MIN_COUNT_NODES_FOR_EXPIRY,
		Dll:                   NewDLL[K, V](),
	}
}
