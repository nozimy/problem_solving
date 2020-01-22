package problems

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func RunReverseLinkedListProblem() {
	head := &Node{Data: 1}
	head.Next = &Node{Data: 2}
	head.Next.Next = &Node{Data: 3}
	head.Next.Next.Next = &Node{Data: 4}

	head = reverse(head)

	fmt.Printf("Reversed Linked List: \n")
	printList(head)
}

func reverse(head *Node) *Node {
	prev := head
	current := head.Next
	var next *Node

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	head.Next = nil
	head = prev

	return head
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%+v\n", head)

		head = head.Next
	}
}
