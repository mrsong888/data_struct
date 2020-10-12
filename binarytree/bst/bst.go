/*
二叉搜索树 插入 删除 查询 前序遍历 中序遍历 后序遍历 递归和非递归
*/
package bst

import (
	"strconv"
	"strings"
	"fmt"
)

type Node struct {
	val int
	left *Node
	right *Node
}

type BST struct {
	root *Node
}

//FIFO
type Queue []*Node

//LIFO
type Stack []*Node

func (q *Queue) Push(node *Node) {
	*q = append(*q, node)
}

func (q *Queue) Pop() *Node {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (s *Stack) Push(node *Node) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *Node {
	lastIndex := len(*s) - 1
	last := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return last
}

func (s *Stack) Front() *Node {
	lastIndex := len(*s) - 1
	last := (*s)[lastIndex]
	return last
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (bst *BST) Insert(val int) {
	newNode := &Node{val:val}
	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
}

func insertNode(root, newNode *Node)  {
	if newNode.val < root.val {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}
}

func (bst *BST) Remove(val int) (*Node, bool) {
	return remove(bst.root, val)
}

func remove(root *Node, val int) (*Node, bool) {
	if root == nil {
		return nil, false
	}
	var existed bool
	if val < root.val {
		root.left, existed = remove(root.left, val)
		return root, existed
	}
	if val > root.val {
		root.right, existed = remove(root.right, val)
		return root, existed
	}
	existed = true
	//leaf node
	if root.left == nil && root.right == nil {
		root = nil
		return root, existed
	}
	//right node
	if root.left == nil {
		root = root.right
		return root, existed
	}
	//left node
	if root.right == nil {
		root = root.left
		return root, existed
	}
	smallestInRight, _ := min(root.right)
	root.val = smallestInRight
	root.right, _ = remove(root.right, smallestInRight)
	return root, existed
}

func (bst *BST) Min() (int, bool) {
	return min(bst.root)
}

func min(root *Node) (int, bool) {
	if root == nil {
		return 0, false
	}
	cur := root
	for {
		if cur.left == nil {
			return cur.val, true
		}
		cur = cur.left
	}
}

func (bst *BST) Max() (int, bool) {
	return max(bst.root)
}

func max(root *Node) (int, bool) {
	if root == nil {
		return 0, false
	}
	cur := root
	for {
		if cur.right == nil {
			return cur.val, true
		}
		cur = cur.right
	}
}

func (bst *BST) Search(val int) bool {
	return search(bst.root, val)
}

func search(root *Node, val int) bool {
	if root == nil {
		return false
	}
	if val < root.val {
		return search(root.left, val)
	}
	if val > root.val {
		return search(root.right, val)
	}
	return true
}

func (bst *BST) SearchTwo(val int) bool {
	return searchTwo(bst.root, val)
}

func searchTwo(root *Node, val int) bool {
	for root != nil {
		if root.val == val {
			return true
		} else if root.val > val {
			root = root.left
		} else if root.val < val {
			root = root.right
		}
	}
	return false
}

func (bst BST) String() string {
	if bst.root != nil {
		var q Queue
		var l = make([]string, 0)
		q.Push(bst.root)
		for !q.IsEmpty() {
			node := q.Pop()
			l = append(l, strconv.Itoa(node.val))
			if node.left != nil {
				q.Push(node.left)
			}
			if node.right != nil {
				q.Push(node.right)
			}
		}
		return strings.Join(l, " ")
	}
	return " "
}

func visit(data int) {
	fmt.Print(data)
}

//DLR
func (bst *BST) PreOrder() {
	preOrder(bst.root)
}

func preOrder(root *Node) {
	if root != nil {
		visit(root.val)
		preOrder(root.left)
		preOrder(root.right)
	}
}

func (bst *BST) PreOrderTwo() {
	preOrderTwo(bst.root)
}

func preOrderTwo(root *Node) {
	var s Stack
	for root != nil || !s.IsEmpty() {
		for root != nil {
			visit(root.val)
			s.Push(root)
			root = root.left
		}
		if !s.IsEmpty() {
			root = s.Pop()
			root = root.right
		}
	}
}

//LDR
func (bst *BST) InOrder() {
	inOrder(bst.root)
}

func inOrder(root *Node) {
	if root != nil {
		inOrder(root.left)
		visit(root.val)
		inOrder(root.right)
	}
}

func (bst *BST) InOrderTwo() {
	inOrderTwo(bst.root)
}

func inOrderTwo(root *Node) {
	var s Stack
	for root != nil || !s.IsEmpty() {
		for root != nil {
			s.Push(root)
			root = root.left
		}
		if !s.IsEmpty() {
			root = s.Pop()
			visit(root.val)
			root = root.right
		}
	}
}

//LRD
func (bst *BST) PostOrder() {
	postOrder(bst.root)
}

func postOrder(root *Node) {
	if root != nil {
		postOrder(root.left)
		postOrder(root.right)
		visit(root.val)
	}
}

func (bst *BST) PostOrderTwo() {
	postOrderTwo(bst.root)
}

func postOrderTwo(root *Node) {
	var s Stack
	var q *Node
	for root != nil || !s.IsEmpty() {
		for root != nil {
			s.Push(root)
			root = root.left
		}
		if !s.IsEmpty() {
			root = s.Front()
			if root.right == nil || root.right == q {
				visit(root.val)
				s.Pop()
				q = root
				root = nil
			} else {
				root = root.right
			}
		}
	}
}

func (bst *BST) Root() *Node {
	return bst.root
}

func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	} else {
		lhight := MaxDepth(root.left)
		rhight := MaxDepth(root.right)
		return Max(lhight, rhight) + 1
	}
}

func Max(l, r int) int {
	if l > r {
		return l
	}
	return r
}