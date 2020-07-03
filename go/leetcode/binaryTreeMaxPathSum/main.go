package main

import(
	"math"
)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    if root ==nil{
        return 0
    }
    pathRoot:=pathSum(root)
    pathLeft:=pathSum(root.Left)
    pathRight:=pathSum(root.Right)

    if root.Left == nil{
        pathLeft = math.MinInt32
    }
    if root.Right == nil{
        pathRight == math.MinInt32
    }

    maxer:=max(root.Val, pathRoot, pathLeft, pathRight)

    if maxer == pathLeft{
        return maxPathSum(root.Left) 
    }else if maxer == pathRight{
        return  maxPathSum(root.Right)
    }else{
        return better(root.Val, pathRoot)
    }
}


func pathSum(node *TreeNode)int{
    if node == nil{
        return 0
    }

    return pathSum(node.Left)+pathSum(node.Right)+node.Val
}

func max(a,b,c,d int)int{
    return better(better(a,b), better(c,d))
}

func better(a,b int)int{
    if a>b{
        return a
    }
    return b
}



func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    var maxGain func(*TreeNode) int
    maxGain = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        // 递归计算左右子节点的最大贡献值
        // 只有在最大贡献值大于 0 时，才会选取对应子节点
        leftGain := max(maxGain(node.Left), 0)
        rightGain := max(maxGain(node.Right), 0)

        // 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
        priceNewPath := node.Val + leftGain + rightGain

        // 更新答案
        maxSum = max(maxSum, priceNewPath)

        // 返回节点的最大贡献值
        return node.Val + max(leftGain, rightGain)
    }
    maxGain(root)
    return maxSum
}
