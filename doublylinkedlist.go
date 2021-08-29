package datastructures

import "fmt"

func NewDoublyLinkedList() doublyLinkedList {
	return doublyLinkedList{}
}

func (dl *doublyLinkedList) Postpend(entry int) {

	node := &nodeD{entry: entry}
	if dl.head == nil {
		dl.head = node
		dl.tail = node
		dl.count++
		return
	}

	dl.tail.nextNode = node
	node.previousNode = dl.tail
	dl.tail = node
	dl.count++

}

func (dl *doublyLinkedList) Prepend(entry int) {

	node := &nodeD{entry: entry}
	if dl.head == nil {
		dl.head = node
		dl.count++
		return
	}

	node.nextNode = dl.head
	dl.head.previousNode = node
	dl.head = node
	dl.count++
}

func (dl *doublyLinkedList) GetAtIndex(index int) (int, bool) {

	if dl.Empty() {
		fmt.Println("\nEmpty List")
		return 0, false
	}

	if index < 0 || index >= dl.count+1 {
		fmt.Printf("\nIndex %v is invalid\n", index)
		return 0, false
	}

	node := dl.head
	i := 0
	for i < index {
		i++
		node = node.nextNode
	}

	return node.entry, true
}

func (dl *doublyLinkedList) InsertAtIndex(index, entry int) {

	if dl.Empty() {
		fmt.Println("\nEmpty List")
		return
	}

	if index < 0 || index >= dl.count+1 {
		fmt.Printf("\nIndex %v is invalid\n", index)
		return
	}

	newNode := &nodeD{entry: entry}
	i := 0
	node := dl.head

	if index == 0 { // Same as prepend
		newNode.nextNode = dl.head
		dl.head.previousNode = newNode
		dl.head = newNode
		dl.count++
		return
	}

	for i < index-1 {
		i++
		node = node.nextNode
	}

	if node.nextNode == nil { // Same as postpend
		node.nextNode = newNode
		newNode.previousNode = node
		dl.count++
		return
	}

	node.nextNode.previousNode = newNode
	newNode.nextNode = node.nextNode
	newNode.previousNode = node
	node.nextNode = newNode
	dl.count++
}

func (dl *doublyLinkedList) DeleteEntry(entry int) {

	if dl.Empty() {
		fmt.Println("\nEmpty List")
		return
	}

	node := dl.head

	for node != nil && node.entry != entry {
		node = node.nextNode
	}

	if node == nil {
		fmt.Printf("\nEntry %v not in List\n", entry)
		return
	}

	if node.previousNode == nil { // Only one nodeD
		dl.head = nil
		dl.count--
		return
	}

	node.previousNode.nextNode = node.nextNode
	dl.count--
}

func (dl *doublyLinkedList) Clear() {
	if dl.Empty() {
		fmt.Println("Lista Vazia")
		return
	}

	dl.head = nil
	dl.count = 0
}

func (dl *doublyLinkedList) Empty() bool {
	return dl.count == 0
}

func (dl *doublyLinkedList) Size() int {
	return dl.count
}

func (dl doublyLinkedList) PrintAllValues() {
	if dl.Empty() {
		fmt.Println("\nEmpty List")
		return
	}
	node := dl.head
	for dl.count != 0 {
		fmt.Printf("%v ", node.entry)
		node = node.nextNode
		dl.count--
	}
	fmt.Println()
}

func (dl doublyLinkedList) ReversePrintAllValues() {
	if dl.Empty() {
		fmt.Println("\nEmpty List")
		return
	}
	node := dl.tail
	for dl.count != 0 {
		fmt.Printf("%v ", node.entry)
		node = node.previousNode
		dl.count--
	}
	fmt.Println()
}

type nodeD struct {
	nextNode     *nodeD
	previousNode *nodeD
	entry        int
}

type doublyLinkedList struct {
	count int
	head  *nodeD
	tail  *nodeD
}
