package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type Project struct {
	profit  int
	capital int
}

type Projects []Project

func (h Projects) Len() int { return len(h) }
func (h Projects) Less(i, j int) bool {
	if h[i].profit == h[j].profit {
		return h[i].capital < h[j].capital
	}
	return h[i].profit > h[j].profit
}
func (h Projects) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Projects) Push(x interface{}) {
	*h = append(*h, x.(Project))
}

func (h *Projects) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

const (
	defaultLength = 10000
)

type Stack struct {
	top     int
	size    int
	element []interface{}
}

/*
*

	根据给定的大小初始话stack
*/
func NewStackBySize(size int) *Stack {
	return &Stack{size, size, make([]interface{}, size)}
}

/*
*

	根据默认的大小初始话stack
*/
func NewStack() *Stack {
	return NewStackBySize(defaultLength)
}

/*
*

	判断stack是否为空
*/
func (stack *Stack) IsEmpty() bool {
	return stack.top == stack.size
}

/*
*

	判断stack是否已经满了
*/
func (stack *Stack) IsFull() bool {
	return stack.top == 0
}

/*
*

	清空stack
*/
func (stack *Stack) Clear() {
	stack.top = stack.size
}

/*
*

	弹出一个元素
*/
func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() == true {
		return nil, errors.New("The Stack is empty")
	}
	stack.top = stack.top + 1
	return stack.element[stack.top-1], nil
}

/*
*

	压入一个元素
*/
func (stack *Stack) Push(e interface{}) error {
	if stack.IsFull() == true {
		return errors.New("The Stack is full")
	}
	stack.top = stack.top - 1
	stack.element[stack.top] = e
	return nil
}
func (stack *Stack) PrintStack() {
	i := stack.top
	for {
		if i == stack.size {
			break
		}
		fmt.Print(stack.element[i], "\n")
		i = i + 1
	}
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := make(Projects, 0, len(profits))
	for i := 0; i < len(profits); i++ {
		projects = append(projects, Project{profit: profits[i], capital: capital[i]})
	}
	heap.Init(&projects)
	s := NewStack()
	surplus := 0
	for i := 0; i < k; i++ {
		if projects.Len() == 0 {
			break
		}
		p := heap.Pop(&projects).(Project)
		if p.capital <= w {
			surplus = k - i
			heap.Push(&projects, p)
			break
		}
		_ = s.Push(p)
		for projects.Len() > 0 {
			p = heap.Pop(&projects).(Project)
			if p.capital <= w {
				w += p.profit
				break
			}
			_ = s.Push(p)
		}
		for !s.IsEmpty() {
			o, _ := s.Pop()
			heap.Push(&projects, o)
		}
	}
	for i := 0; i < surplus; i++ {
		w += heap.Pop(&projects).(Project).profit
	}
	return w
}

func main() {
	//fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
	fmt.Println(findMaximizedCapital(1, 2, []int{1, 2, 3}, []int{1, 1, 2}))
}
