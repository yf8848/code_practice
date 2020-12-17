package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param pRoot TreeNode类
 * @return int整型二维数组
 */
func Print(pRoot *TreeNode) [][]int {
	// write code here
	queue := [][]*TreeNode{}
	queue = append(queue, []*TreeNode{pRoot})
	res := [][]int{}

	for len(queue) > 0 {
		row := []int{}
		rowNode := []*TreeNode{}

		nodes := queue[0]
		queue = queue[1:]

		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			if node == nil {
				continue
			}

			if node.Left != nil {
				rowNode = append(rowNode, node.Left)
			}

			if node.Right != nil {
				rowNode = append(rowNode, node.Right)
			}

			row = append(row, node.Val)
		}

		res = append(res, row)
		if len(rowNode) > 0 {
			queue = append(queue, rowNode)
		}
	}
	return res
}

func createTree() *TreeNode {
	root := &TreeNode{
		Val: 1,
	}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	return root
}

func main() {
	fmt.Println(Print(createTree()))
}
