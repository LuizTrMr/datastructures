package datastructures

import "fmt"

func NewBinarySearchTree() binarySearchTree {
	return binarySearchTree{}
}

// Iterative Insert
func (bst *binarySearchTree) Insert(entry int) {
	tNode := bst.root
	var pNode *treeNode = nil
	for tNode != nil {
		pNode = tNode
		if entry > tNode.entry {
			tNode = tNode.rightNode
		} else if entry < tNode.entry {
			tNode = tNode.leftNode
		}
	}
	newTNode := &treeNode{entry: entry}
	if pNode == nil {
		bst.root = newTNode
	} else {
		if entry < pNode.entry {
			pNode.leftNode = newTNode
		} else {
			pNode.rightNode = newTNode
		}
	}
}

// Recursive Insert
func (bst *binarySearchTree) InsertR(entry int) {
	bst.root = bst.insertR(entry, bst.root)
}

func (bst *binarySearchTree) insertR(entry int, ptreeNode *treeNode) *treeNode {
	if ptreeNode == nil {
		ptreeNode = &treeNode{entry: entry}
		return ptreeNode
	}

	if entry > ptreeNode.entry {
		ptreeNode.rightNode = bst.insertR(entry, ptreeNode.rightNode)
	} else if entry < ptreeNode.entry {
		ptreeNode.leftNode = bst.insertR(entry, ptreeNode.leftNode)
	}
	return ptreeNode
}

// Iterative Search
func (bst *binarySearchTree) Search(entry int) bool {
	tNode := bst.root
	var inTree bool = false
	for tNode != nil && tNode.entry != entry {
		if entry > tNode.entry {
			tNode = tNode.rightNode
		} else if entry < tNode.entry {
			tNode = tNode.leftNode
		}
	}
	if tNode != nil {
		inTree = true
	}
	return inTree
}

// Recursive Search
func (bst *binarySearchTree) SearchR(entry int) bool {
	return bst.searchR(entry, bst.root)
}

func (bst *binarySearchTree) searchR(entry int, ptreeNode *treeNode) bool {

	// Base Case
	if ptreeNode == nil {
		return false
	}

	// Base Case
	if ptreeNode.entry == entry {
		return true
	}

	if entry > ptreeNode.entry {
		return bst.searchR(entry, ptreeNode.rightNode)
	} else {
		return bst.searchR(entry, ptreeNode.leftNode)
	}
}

// Iterative Remove
func (bst *binarySearchTree) Remove(entry int) {
	bst.root = bst.remove(entry, bst.root)
}

func (bst *binarySearchTree) remove(entry int, ptreeNode *treeNode) *treeNode {
	currNode := ptreeNode
	var prevNode *treeNode = nil

	for currNode != nil && currNode.entry != entry {
		prevNode = currNode

		if entry > currNode.entry {
			currNode = currNode.rightNode
		} else {
			currNode = currNode.leftNode
		}
	}

	if currNode == nil {
		fmt.Println("\nEntry not in Tree")
		return nil
	}

	newCurr := &treeNode{}

	if currNode.rightNode == nil || currNode.leftNode == nil {

		if currNode.rightNode == nil {
			if currNode.leftNode == nil { // Node has 0 children
				newCurr = nil
			}
			newCurr = currNode.leftNode // Node has one child, in its left
		} else if currNode.leftNode == nil {
			newCurr = currNode.rightNode
		}

		if prevNode == nil { // Node to be removed is in the root
			return newCurr
		}

		if currNode == prevNode.leftNode {
			prevNode.leftNode = newCurr
		} else {
			prevNode.rightNode = newCurr
		}

	} else { // Node has 2 children

		var aux *treeNode = nil
		temp := currNode.rightNode
		for temp.leftNode != nil {
			aux = temp
			temp = temp.leftNode
		}

		if aux != nil {
			aux.leftNode = temp.rightNode
		} else {
			currNode.rightNode = temp.rightNode
		}
		currNode.entry = temp.entry
	}

	return ptreeNode
}

// Recursive Remove
func (bst *binarySearchTree) RemoveR(entry int) {
	bst.root = bst.removeR(entry, bst.root)

}

func (bst *binarySearchTree) removeR(entry int, ptreeNode *treeNode) *treeNode {
	if ptreeNode == nil {
		return ptreeNode
	}

	if entry > ptreeNode.entry {
		ptreeNode.rightNode = bst.removeR(entry, ptreeNode.rightNode)
	} else if entry < ptreeNode.entry {
		ptreeNode.leftNode = bst.removeR(entry, ptreeNode.leftNode)
	} else {

		if ptreeNode.rightNode == nil {
			if ptreeNode.leftNode == nil { // Node has 0 children
				return nil
			}
			return ptreeNode.leftNode
		} else if ptreeNode.leftNode == nil {
			return ptreeNode.rightNode
		}

		// Node has 2 children
		ptreeNode.entry = bst.min(ptreeNode.rightNode)

		ptreeNode.rightNode = bst.removeR(ptreeNode.entry, ptreeNode.rightNode)
	}

	return ptreeNode
}

func (bst *binarySearchTree) Clear() {
	bst.root = nil
}

func (bst *binarySearchTree) Empty() bool {
	return bst.root == nil
}

func (bst *binarySearchTree) Size() int {
	if bst.Empty() {
		return 0
	}
	return bst.size(bst.root)
}

func (bst *binarySearchTree) size(ptreeNode *treeNode) int {
	leftSize, rightSize := 0, 0

	if ptreeNode.leftNode != nil {
		leftSize = bst.size(ptreeNode.leftNode)
	}

	if ptreeNode.rightNode != nil {
		rightSize = bst.size(ptreeNode.rightNode)
	}

	return 1 + leftSize + rightSize
}

func (bst *binarySearchTree) InOrder() {
	if bst.Empty() {
		fmt.Println("\nEmpty Tree")
		return
	}
	fmt.Print("\nBinary Search Tree In Order : ")
	inOrder(bst.root)
}

func inOrder(tNode *treeNode) {
	if tNode == nil {
		return
	}
	inOrder(tNode.leftNode)
	fmt.Printf("%d ", tNode.entry)
	inOrder(tNode.rightNode)
}

func (bst *binarySearchTree) min(ptreeNode *treeNode) int {
	min := ptreeNode.entry
	for ptreeNode.leftNode != nil {
		min = ptreeNode.leftNode.entry
		ptreeNode = ptreeNode.leftNode
	}
	return min
}

type treeNode struct {
	entry     int
	leftNode  *treeNode
	rightNode *treeNode
}

type binarySearchTree struct {
	root *treeNode
}
