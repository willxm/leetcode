package problems

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}

//用双向链表实现stack
type Element interface{}

var header *entry //链表表头
var size int      //栈的长度

type entry struct {
	previous *entry
	next     *entry
	element  Element
}

func newEntry(prev, next *entry, e Element) *entry {
	return &entry{prev, next, e}
}

//初始化header  表头
func NewStack() *entry {
	header = newEntry(nil, nil, nil)
	header.previous = header
	header.next = header
	return header
}

type Stack interface {
	Push(e Element) //向栈顶添加元素
	Pop() Element   //移除栈顶元素
	Top() Element   //获取栈顶元素（不删除）
	Clear() bool    //清空栈
	Size() int      //获取栈的元素个数
	IsEmpty() bool  //判断栈是否是空栈
}

//向栈顶添加元素
func (e *entry) Push(element Element) {
	addBefore(header, element)
}

//移除栈顶元素
func (e *entry) Pop() Element {
	if e.IsEmpty() {
		fmt.Println("stack is empty!")
		return nil
	}
	prevEntry := header.previous

	prevEntry.previous.next = header
	header.previous = prevEntry.previous

	size--
	return prevEntry.element
}

//获取栈顶元素（不删除）
func (e *entry) Top() Element {
	if e.IsEmpty() {
		fmt.Println("stack is empty!")
		return nil
	}
	return header.previous.element
}

func (e *entry) Size() int {
	return size
}

func (e *entry) IsEmpty() bool {
	if size == 0 {
		return true
	}

	return false
}

//在entry节点之前添加
func addBefore(e *entry, element Element) Element {
	newEntry := newEntry(e.previous, e, element)
	newEntry.previous.next = newEntry
	newEntry.next.previous = newEntry
	size++
	return newEntry
}

func inorderTraversal1(root *TreeNode) []int {
	p := root
	s := NewStack()

	var res []int

	for !s.IsEmpty() || p != nil {
		if p != nil {
			s.Push(p)
			p = p.Left
		} else {
			top := s.Top().(*TreeNode)

			s.Pop()
			res = append(res, top.Val)
			p = top.Right
		}
	}
	return res
}
