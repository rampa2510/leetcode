package main

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

type doubleLinkedList struct {
	head *node
}

func createNewNode(data int) *node {
	var newNode *node
	newNode = new(node)
	newNode.data = data
	newNode.prev, newNode.next = nil, nil
	return newNode
}

func (ll *doubleLinkedList) isEmpty() bool {
	if ll.head == nil {
		return true
	}
	return false
}

func (ll *doubleLinkedList) insertAtBeg(data int) {
	var newNode *node
	newNode = createNewNode(data)
	if ll.isEmpty() {
		newNode.next, ll.head = ll.head, newNode
		return
	}
	ll.head.prev = newNode
	newNode.next = ll.head
	ll.head = newNode
	// fmt.Println(newNode, ll.head)
}
func (ll *doubleLinkedList) insert(data, pos int) {
	var newNode *node
	newNode = createNewNode(data)
	if ll.isEmpty() && pos > 1 {
		fmt.Println("Empty linked list so will be inserting at start")
		ll.head = newNode
		return
	}
	currNode := ll.head
	for i := 1; i < pos-1; i++ {
		currNode = currNode.next
	}
	fmt.Println(currNode)
	newNode.next, newNode.prev, currNode.next.prev, currNode.next = currNode.next, currNode, newNode, newNode
	fmt.Println(newNode)
}
func (ll *doubleLinkedList) insertAtEnd(data int) {
	var newNode *node
	newNode = createNewNode(data)
	if ll.isEmpty() {
		fmt.Println("Empty linked list")
		return
	}
	currNode := ll.head
	for currNode.next != nil {
		currNode = currNode.next
	}
	currNode.next, newNode.prev = newNode, currNode
}

func (ll *doubleLinkedList) displayFromStart() {
	if ll.isEmpty() {
		fmt.Println("Empty linked list")
		return
	}
	currNode := ll.head
	for currNode != nil {
		fmt.Println(currNode.data)
		currNode = currNode.next
	}
}

func (ll *doubleLinkedList) displayFromStartWRec(currNode *node) {
	if ll.isEmpty() {
		fmt.Println("Empty linked list")
		return
	}
	if currNode == nil {
		return
	}
	fmt.Println(currNode.data)
	ll.displayFromStartWRec(currNode.next)
}

func (ll *doubleLinkedList) displayFromEnd() {
	currNode := ll.head
	for currNode.next != nil {
		currNode = currNode.next
	}
	// fmt.Println(currNode)
	for currNode != nil {
		fmt.Println(currNode.data)
		currNode = currNode.prev
	}
}

func main() {
	ll := &doubleLinkedList{head: nil}
	ll.insertAtBeg(2)
	ll.insertAtBeg(3)
	ll.insertAtBeg(4)
	ll.insertAtBeg(5)
	ll.displayFromStart()
	fmt.Println("Display after inserting at end")
	ll.insertAtEnd(6)
	ll.insertAtEnd(7)
	ll.insertAtEnd(8)
	ll.displayFromStart()
	fmt.Println("Display after inserting at random")
	ll.insert(9, 2)
	ll.displayFromStart()
	fmt.Println("Print from end")
	ll.displayFromEnd()
}
