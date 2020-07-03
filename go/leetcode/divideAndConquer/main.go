package main

/**
 * 分治模板
 * - 递归返回条件
 * - 分段处理
 * - 合并结果
 */

/**
 * 递归遍历二叉树，自低向上
 */
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/**
 * DFS
 * 自顶向下
 */
func orderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)
	dfs(node.Left, res)
	dfs(node.Right, res)
}

/**
 * DFS
 * 自底向上(分治法)
 */
func preOrderTravlSal(root *TreeNode) []int {
	return divideAndConquer(root)
}

func divideAndConquer(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)
	return res
}

/**
 * merge sort
 */

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(a, b []int) []int {
	l := 0
	r := 0
	res := make([]int, 0, len(b)+len(a))

	i := 0
	for l < len(a) && l < len(b) {
		if a[i] < b[i] {
			res = append(res, a[i])
			l++

		} else {
			res = append(res, b[i])
			r++
		}
	}

	res = append(res, a[l:]...)
	res = append(res, b[r:]...)
	return res
}

/**
 * quick sort
 */
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		provit := partition(nums, start, end)
		quickSort(nums, start, provit-1)
		quickSort(nums, provit+1, end)
	}
}

func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, j, i)
			i++
		}
	}
	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

/**
 * 左右子树最近公共祖先
 */

func lowestCommonAncestorOfBT(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestorOfBT(root.Left, p, q)
	right := lowestCommonAncestorOfBT(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}

	if right != nil {
		return right
	}

	return nil
}

/**
 * BFS
 * 层序遍历
 */
func lelvelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]

			list = append(list, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, list)
	}
	return res
}

/**
 * BFS
 * zig-zag z 形层序遍历
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	toggle := false
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)

		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]

			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if toggle {
			resverse(list)
		}
		res = append(res, list)
		toggle = !toggle
	}
	return res
}

func resverse(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

/**
 * 判断二叉搜索树
 * - 中序遍历，判断有序
 * - 判断左 max<root.val < 右 min
 */

func isVaildBT(root *TreeNode) bool {
	//v1: 中序遍历+判断有序
	res := make([]int, 0)
	inOrderTraverSal(root, &res)
	return isOrder(res)

	//v2: 左 max<root.Val<右 min
	return isVaildBST(root)
}

func inOrderTraverSal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inOrderTraverSal(root.Left, res)
	*res = append(*res, root.Val)
	inOrderTraverSal(root.Right, res)
}

func isOrder(list []int) bool {
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}

/**
 * 分治：判断二叉平衡树
 */
type resultType struct {
	Res      bool
	Max, Min *TreeNode
}

func isVaildBST(root *TreeNode) bool {
	res := helper(root)
	return res.Res
}

func helper(root *TreeNode) resultType {
	res := resultType{}
	if root == nil {
		res.Res = true
		return res
	}

	left := helper(root.Left)
	right := helper(root.Right)

	if !left.Res || !right.Res {
		res.Res = false
		return res
	}

	if left.Max != nil && left.Max.Val >= root.Val {
		res.Res = false
		return res
	}

	if right.Min != nil && right.Min.Val <= root.Val {
		res.Res = false
		return res
	}
	res.Res = true

	res.Min = root
	if left.Min != nil {
		res.Min = left.Min
	}

	res.Max = root
	if right.Max != nil {
		res.Max = right.Max
	}
	return res
}

func insertIntoBst(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}

	if root.Val > val {
		root.Left = insertIntoBst(root.Left, val)
	} else {
		root.Right = insertIntoBst(root.Right, val)
	}
	return root
}
