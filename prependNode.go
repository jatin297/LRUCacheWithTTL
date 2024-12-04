package main

func (dll *DoublyLinkedList[K, V]) Prepend(node *Node[K, V]) {
	// check if dll is empty
	if dll.Head == nil {
		dll.Head, dll.Tail = node, node
		return
	}

	node.Next = dll.Head
	node.Prev = nil
	dll.Head.Prev = node
	dll.Head = node
}
