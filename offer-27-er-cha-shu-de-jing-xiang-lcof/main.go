package main

import (
	. "github.com/xjian2021/leetcode/pkg"
)

//剑指 Offer 27. 二叉树的镜像
//请完成一个函数，输入一个二叉树，该函数输出它的镜像。
/*
例如输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9

镜像输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1

示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
*/

func main() {
	A := SliceToTreeNode([]int{4, 2, 4, 8, NULL, NULL, 9, NULL, NULL, 5, NULL, NULL, 3, 6, NULL, NULL, 7})
	A.EchoTreeNode()
	B := mirrorTree2(A)
	B.EchoTreeNode()
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//mirrorTree BFS迭代
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		if q[0].Left != nil {
			q = append(q, q[0].Left)
		}
		if q[0].Right != nil {
			q = append(q, q[0].Right)
		}
		q[0].Left, q[0].Right = q[0].Right, q[0].Left
		q = q[1:]
	}
	return root
}

//mirrorTree2 DFS递归
func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Right, root.Left = mirrorTree2(root.Left), mirrorTree2(root.Right)
	return root
}
