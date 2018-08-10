package main

import "fmt"

func main() {
	ll := LinkedList{}

	ll.InsertBefore(Node{Value: 100}, Node{Value: 1})
	ll.InsertLast(Node{Value: 200})
	ll.InsertFirst(Node{Value: 100})

	ll.Insert(Node{Value: 1})
	ll.Insert(Node{Value: 2})
	ll.Insert(Node{Value: 3})
	ll.Insert(Node{Value: 4})
	printList(ll)

	ll.Reverse()
	printList(ll)

	ll.InsertFirst(Node{Value: 5})
	printList(ll)

	ll.InsertBefore(Node{Value: 11}, Node{Value: 3})
	printList(ll)

	ll.InsertLast(Node{Value: 10})
	printList(ll)

	ll.Remove(Node{Value: 1})
	ll.Remove(Node{Value: 5})
	ll.Remove(Node{Value: 10})
	ll.RemoveFirst()
	ll.RemoveLast()
	printList(ll)

	ll.Reverse()
	printList(ll)
}

func printList(ll LinkedList) {
	current := ll.Head()
	for current != nil {
		fmt.Printf("{%p}Node: {%v}\n", current, current)
		current = current.Next
	}
	fmt.Printf("Size: %d", ll.Size())
	fmt.Print("\n\n")
}
