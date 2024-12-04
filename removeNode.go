package main

func (dll *DoublyLinkedList[K, V]) Remove(node *Node[K, V]) {

	if node == dll.Head {
		dll.Head = dll.Head.Next
	}

	if node == dll.Tail {
		dll.Tail = dll.Tail.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Prev = nil
	node.Next = nil
}
