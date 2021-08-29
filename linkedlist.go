package datastructures

import "fmt"

func LinkedList() {
	l := NewLinkedList()

	l.Postpend(4)
	l.InsertAtIndex(2, 89)
	l.PrintAllValues()
}

func NewLinkedList() linkedList {
	return linkedList{}
}

func (ll *linkedList) Prepend(entry int) {
	newNode := &node{entry: entry}
	aux := ll.head
	ll.head = newNode
	newNode.nextNode = aux
	ll.count++
}

func (ll *linkedList) Postpend(entry int) {
	newNode := &node{entry: entry}

	if ll.Empty() {
		ll.head = newNode
		ll.count++
		return
	}

	node := ll.head
	for node.nextNode != nil {
		node = node.nextNode
	}

	node.nextNode = newNode
	ll.count++
}

func (ll *linkedList) GetAtIndex(index int) (int, bool) {

	if ll.Empty() {
		fmt.Println("\nEmpty List")
		return 0, false
	}

	if index < 0 || index >= ll.count+1 {
		fmt.Printf("\nIndex %v is invalid\n", index)
		return 0, false
	}

	node := ll.head
	i := 0
	for i < index {
		i++
		node = node.nextNode
	}

	return node.entry, true
}

func (ll *linkedList) InsertAtIndex(index, entry int) { // Inserting at (index == ll.count) is the same as postpending

	if ll.Empty() {
		fmt.Println("\nEmpty List")
		return
	}

	if index < 0 || index >= ll.count+1 {
		fmt.Printf("\nIndex %v is invalid\n", index)
		return
	}

	newNode := &node{entry: entry}
	i := 0
	node := ll.head

	if index == 0 {
		newNode.nextNode = ll.head
		ll.head = newNode
		ll.count++
		return
	}

	for i < index-1 {
		i++
		node = node.nextNode
	}

	newNode.nextNode = node.nextNode
	node.nextNode = newNode
	ll.count++
}

func (ll *linkedList) DeleteEntry(entry int) {
	if ll.Empty() {
		fmt.Println("\nEmpty List")
		return
	}

	if ll.head.entry == entry {
		ll.head = ll.head.nextNode
		ll.count--
		return
	}

	preToDelete := ll.head
	for preToDelete.nextNode != nil && preToDelete.nextNode.entry != entry {
		preToDelete = preToDelete.nextNode
	}

	if preToDelete.nextNode == nil {
		fmt.Printf("\nEntry %v not in List\n", entry)
		return
	}

	preToDelete.nextNode = preToDelete.nextNode.nextNode
	ll.count--
}

func (ll *linkedList) Clear() {

	if ll.Empty() {
		fmt.Println("\nEmpty List")
		return
	}

	ll.head = nil
	ll.count = 0
}

func (ll *linkedList) Empty() bool {
	return ll.count == 0
}

func (ll *linkedList) Size() int {
	return ll.count
}

func (ll linkedList) PrintAllValues() {
	if ll.Empty() {
		fmt.Println("\nEmpty List")
		return
	}
	fmt.Print("\nFull list : ")
	node := ll.head
	for ll.count != 0 {
		fmt.Printf("%v ", node.entry)
		node = node.nextNode
		ll.count--
	}
	fmt.Println()
}

type node struct {
	nextNode *node
	entry    int
}

type linkedList struct {
	count int
	head  *node
}
