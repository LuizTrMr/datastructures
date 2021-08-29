package datastructures

import "fmt"

func NewMaxHeap() heap {
	return heap{max: true}
}

func NewMinHeap() heap {
	return heap{max: false}
}

func (h *heap) Insert(entry int) {
	h.array = append(h.array, entry)
	h.heapifyUp()
}

func (h *heap) Extract() (int, bool) {

	if h.Empty() {
		fmt.Println("\nEmpty Heap")
		return 0, false
	}

	toReturn := h.array[0]
	h.swap(0, len(h.array)-1)
	h.array = h.array[:len(h.array)-1]

	h.heapifyDown()

	return toReturn, true
}

func (h *heap) heapifyUp() {
	indexOfEntry := (len(h.array) - 1)

	if h.max {
		for hasParent(indexOfEntry) && h.array[indexOfEntry] > h.array[parentIndex(indexOfEntry)] { // For max heap
			h.swap(indexOfEntry, parentIndex(indexOfEntry))
			indexOfEntry = parentIndex(indexOfEntry)
		}
	} else {
		for hasParent(indexOfEntry) && h.array[indexOfEntry] < h.array[parentIndex(indexOfEntry)] { // For min heap
			h.swap(indexOfEntry, parentIndex(indexOfEntry))
			indexOfEntry = parentIndex(indexOfEntry)
		}
	}
}

func (h *heap) heapifyDown() {
	indexOfEntry := 0

	lastIndex := len(h.array) - 1
	lIndex := leftIndex(indexOfEntry)
	rIndex := rightIndex(indexOfEntry)

	if h.max {
		var indexOfMaxChild int
		for lIndex <= lastIndex {
			if lIndex == lastIndex || h.array[lIndex] > h.array[rIndex] { // Item only has one child or left child has the highest value
				indexOfMaxChild = lIndex
			} else {
				indexOfMaxChild = rIndex
			}

			if h.array[indexOfMaxChild] > h.array[indexOfEntry] {
				h.swap(indexOfEntry, indexOfMaxChild)
				indexOfEntry = indexOfMaxChild
				lIndex = leftIndex(indexOfEntry)
				rIndex = rightIndex(indexOfEntry)
			} else {
				break
			}
		}
	} else {
		var indexOfMinChild int
		for lIndex <= lastIndex {
			if lIndex == lastIndex || h.array[lIndex] < h.array[rIndex] { // Item only has one child or left child has the lowest value
				indexOfMinChild = lIndex
			} else {
				indexOfMinChild = rIndex
			}

			if h.array[indexOfMinChild] < h.array[indexOfEntry] {
				h.swap(indexOfEntry, indexOfMinChild)
				indexOfEntry = indexOfMinChild
				lIndex = leftIndex(indexOfEntry)
				rIndex = rightIndex(indexOfEntry)
			} else {
				break
			}
		}
	}
}

func (h *heap) Peek() (int, bool) {
	if h.Empty() {
		fmt.Println("\nEmpty Heap")
		return 0, false
	} else {
		return h.array[0], true
	}
}

func (h *heap) Clear() {
	h.array = []int{}
}

func (h *heap) Empty() bool {
	return len(h.array) == 0
}

func (h *heap) Size() int {
	return len(h.array)
}

func (h *heap) PrintAllValues() {
	if h.Empty() {
		fmt.Println("\nEmpty Heap")
		return
	}
	if h.max {
		fmt.Print("\nMax Heap : ")
	} else {
		fmt.Print("\nMin Heap : ")
	}
	fmt.Println(h.array)
}

func (h *heap) swap(indexA int, indexB int) {
	aux := h.array[indexA]
	h.array[indexA] = h.array[indexB]
	h.array[indexB] = aux
}

func leftIndex(index int) int {
	return index*2 + 1
}

func rightIndex(index int) int {
	return index*2 + 2
}

func hasParent(index int) bool {
	return parentIndex(index) >= 0
}

func parentIndex(index int) int {
	return (index - 1) / 2
}

type heap struct {
	array []int
	max   bool
}
