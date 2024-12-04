package main

import (
	"time"
)

func main() {

	capacity := 10
	lru := NewLRU[int, interface{}](capacity)

	done := make(chan bool)
	go lru.RunActiveExpirationConcurrently(done)
	defer func() {
		done <- true
		close(done)
	}()

	// putting values in cache
	// key will never expire
	lru.Set(1, 1, -1)

	// key will expire after 0 second
	lru.Set(2, 2, time.Duration(2)*time.Second)

	// key will expire after 1 second
	lru.Set(3, 3, time.Duration(1)*time.Second)
	//
	//// key will expire after 2 second
	lru.Set(4, 4, time.Duration(2)*time.Second)
	//
	//// key will expire after 3 second
	lru.Set(5, 5, time.Duration(3)*time.Second)

	// key will expire after 14 second
	lru.Set(6, 6, time.Duration(7)*time.Second)

	time.Sleep(time.Second * 3)

	lru.Set(7, 8, time.Duration(2)*time.Second)

	lru.Get(7)
	PrintDLL(lru.Dll)
}
