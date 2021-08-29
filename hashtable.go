package datastructures

import (
	"fmt"
)

const count int = 10

func NewHashTable() hashTable {
	return hashTable{}
}

func (h *hashTable) Put(key string, entry int) {
	index := hash(key)

	if h.data[index] == nil { // If a bucket has not yet been created in that index
		h.data[index] = &bucket{head: nil}
	}

	h.data[index].put(key, entry)
}

func (b *bucket) put(key string, entry int) {
	newNode := &bucketNode{key: key, entry: entry}

	if b.head == nil {
		b.head = newNode
	} else {
		node := b.head

		for node.nextBucketNode != nil && node.key != key {
			node = node.nextBucketNode
		}

		if node.key == key {
			node.entry = entry
			return
		}

		node.nextBucketNode = newNode
	}
}

func (h *hashTable) Get(key string) (int, bool) {
	index := hash(key)

	if h.data[index] == nil {
		fmt.Printf("\nKey %v doesn't exist on hash table\n", key)
		return 0, false
	}
	return h.data[index].get(key)
}

func (b *bucket) get(key string) (int, bool) {

	node := &bucketNode{nextBucketNode: b.head}

	for node.nextBucketNode != nil && node.nextBucketNode.key != key {
		node = node.nextBucketNode
	}

	if node.nextBucketNode == nil {
		fmt.Printf("\nKey %v doesn't exist on hash table\n", key)
		return 0, false
	}

	return node.nextBucketNode.entry, true

}

func (h *hashTable) Remove(key string) bool {
	index := hash(key)
	if h.data[index] == nil {
		fmt.Printf("\nKey %v doesn't exist on hash table\n", key)
		return false
	}
	return h.data[index].remove(key)
}

func (b *bucket) remove(key string) bool {

	if b.head == nil {
		fmt.Printf("\nKey %v doesn't exist on hash table\n", key)
		return false
	}

	if b.head.key == key {
		b.head = b.head.nextBucketNode
		return true
	}

	node := b.head

	for node.nextBucketNode != nil && node.nextBucketNode.key != key {
		node = node.nextBucketNode
	}

	if node.nextBucketNode == nil {
		fmt.Printf("\nKey %v doesn't exist on hash table\n", key)
		return false
	}

	node.nextBucketNode = node.nextBucketNode.nextBucketNode
	return true

}

func (h *hashTable) ContainsKey(key string) bool {
	index := hash(key)
	if h.data[index] == nil {
		return false
	}
	return h.data[index].containsKey(key)
}

func (b *bucket) containsKey(key string) bool {
	node := b.head

	for node.nextBucketNode != nil && node.key != key {
		node = node.nextBucketNode
	}

	if node.key == key {
		return true
	} else {
		return false
	}
}

func (h *hashTable) Clear() {
	h.data = [count]*bucket{}
}

func (h *hashTable) Empty() bool {
	for i := range h.data {
		if h.data[i] != nil {
			if h.data[i].head != nil {
				return false
			}
		}
	}
	return true
}

func (h *hashTable) Size() int {

	if h.Empty() {
		return 0
	}

	len := 0
	for i := range h.data {
		if h.data[i] != nil {
			if h.data[i].head != nil {
				node := h.data[i].head
				for node != nil {
					node = node.nextBucketNode
					len++
				}
			}
		}
	}
	return len
}

func (h *hashTable) PrintAllValues() {

	if h.Empty() {
		fmt.Println("\nEmpty Hash Table")
		return
	}

	fmt.Print("\nHash Table : ")
	for i := range h.data {
		if h.data[i] != nil {
			if h.data[i].head != nil {
				node := h.data[i].head
				for node != nil {
					fmt.Printf(" (%v,%v)", node.key, node.entry)
					node = node.nextBucketNode
				}
			}
		}
	}
	fmt.Println()
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return (sum % count)
}

type hashTable struct {
	data [count]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key            string
	entry          int
	nextBucketNode *bucketNode
}
