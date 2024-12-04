package main

import (
	"fmt"
	"time"
)

type Node[K comparable, V any] struct {
	Key   K
	Value V

	// time of creation or last time access
	TimeStamp time.Time

	// ttl (time to live) key expire if current_time > time_stamp + ttl
	TTL time.Duration

	Prev *Node[K, V]
	Next *Node[K, V]
}

type DoublyLinkedList[K comparable, V any] struct {
	Head *Node[K, V]
	Tail *Node[K, V]
}

func NewNode[K comparable, V any](key K, value V, ttl time.Duration) *Node[K, V] {
	return &Node[K, V]{
		Key:       key,
		Value:     value,
		TTL:       ttl,
		TimeStamp: time.Now(),
	}
}

func NewDLL[K comparable, V any]() *DoublyLinkedList[K, V] {
	return &DoublyLinkedList[K, V]{}
}

func PrintDLL[K comparable, V any](dll *DoublyLinkedList[K, V]) {
	curr := dll.Head

	for curr != nil {
		fmt.Print(curr.Key, " <=> ")
		curr = curr.Next
	}

	fmt.Println()
	fmt.Println()

}
func (lru *LRU[K, V]) Print() {
	fmt.Println(lru.Bucket)
	PrintDLL(lru.Dll)
}
