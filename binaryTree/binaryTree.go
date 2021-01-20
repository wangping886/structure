package main

import "fmt"

type binaryNode struct {
	Val   int
	Left  *binaryNode
	Right *binaryNode
}

var tree *binaryNode

func InsertNode(n *binaryNode) {
	if tree == nil {
		tree = n
		return
	}
	curr := tree
	for {
		if curr.Val > n.Val {
			if curr.Left == nil {
				curr.Left = n
				return
			}
			curr = curr.Left
			continue
		} else {
			if curr.Right == nil {
				curr.Right = n
				return
			}
			curr = curr.Right
			continue
		}
		return
	}
}

func traverseTree(t *binaryNode) {
	if t == nil {
		return
	}
	traverseTree(t.Left)
	fmt.Println(t.Val)
	traverseTree(t.Right)
}

//深度优先
func DFSTraverseTree(t *binaryNode) {
	//nodeStack := make([]binaryNode, 0)
	if t == nil {
		return
	}

}

//广度优先对当前每一个节点都应该优先输出他的所有子节点，先访问的先保证输出，可以借助队列
func BFSTraverseTree(t *binaryNode) {
	nodeQueue := make([]*binaryNode, 0)
	if t == nil {
		return
	}
	nodeQueue = append(nodeQueue, t)
	for {
		headQueue := nodeQueue[0]
		if headQueue == nil {
			return
		}
		fmt.Println(headQueue.Val)
		if headQueue.Left != nil {
			nodeQueue = append(nodeQueue, headQueue.Left)
		}
		if headQueue.Right != nil {
			nodeQueue = append(nodeQueue, headQueue.Right)
		}
		if len(nodeQueue) == 1 {
			return
		}
		nodeQueue = nodeQueue[1:]
	}

}

func main() {
	var nums = []int{5, 3, 2, 1, 6, 4, 8}
	for i := 0; i < len(nums); i++ {
		InsertNode(&binaryNode{Val: nums[i]})
	}

	traverseTree(tree)
	fmt.Println(111)

	BFSTraverseTree(tree)

}
