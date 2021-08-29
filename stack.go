package datastructures

import (
	"fmt"
)

const topStart = -1

func NewStack() stackStruct {
	return stackStruct{top: topStart}
}

func (s *stackStruct) GetTop() (int, bool) {
	if s.Empty() {
		fmt.Println("\nEmpty Stack")
		return 0, false
	} else {
		return s.entries[s.top], true
	}
}

func (s *stackStruct) Push(entry int) {
	s.entries = append(s.entries, entry)
	s.top++
}

func (s *stackStruct) Pop() (int, bool) {
	if s.Empty() {
		fmt.Println("\nEmpty Stack")
		return 0, false
	}
	toPop := s.entries[s.top]
	s.entries = s.entries[:s.top]
	s.top--
	return toPop, true
}

func (s *stackStruct) Clear() {
	s.entries = []int{}
	s.top = topStart
}

func (s *stackStruct) Empty() bool {
	return (s.top == topStart)
}

func (s *stackStruct) Size() int {
	return len(s.entries)
}

func (s *stackStruct) PrintAllValues() {
	stack := *s

	if s.Empty() {
		fmt.Println("\nEmpty Stack")
		return
	}

	newStack := stackStruct{top: topStart}

	for !stack.Empty() {
		v, _ := stack.Pop()
		newStack.Push(v)
	}

	str := "Stack :\n"
	for _, v := range newStack.entries {
		str += fmt.Sprint(v) + "\n"
	}

	fmt.Println(str)
}

type stackStruct struct {
	top     int
	entries []int
}
