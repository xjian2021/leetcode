package main

import (
	"fmt"
	. "github.com/xjian2021/leetcode/pkg"
)

//剑指 Offer 28. 对称的二叉树
//请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

/*
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3

示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false
*/

func main() {
	//A := SliceToTreeNode([]int{1, 2, 3, NULL, NULL, 4, NULL, NULL, 2, 4, NULL, NULL, 3})
	A := SliceToTreeNode([]int{1, 2, NULL, 3, NULL, NULL, 2, NULL, 3})
	//A := SliceToTreeNode([]int{4, 2, 4, 8, NULL, NULL, 9, NULL, NULL, 5, NULL, NULL, 3, 6, NULL, NULL, 7})
	A.EchoTreeNode()
	//B := SliceToTreeNode([]int{4, 2, 4, 8, NULL, NULL, 9, NULL, NULL, 5, NULL, NULL, 3, 6, NULL, NULL, 7})
	//B = mirrorTree(B)
	//B.EchoTreeNode()
	fmt.Println(isSymmetric2(A))
	//fmt.Println(isSymmetric2(&TreeNode{
	//	Val:   3,
	//	Left:  A,
	//	Right: B,
	//}))
}

/**
 * Definition for a binary tree node.s
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归对比左右两边每个节点是否相同
// isSymmetric DFS递归
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return symmetric(a.Left, b.Right) && symmetric(a.Right, b.Left)
}

// isSymmetric2 BFS迭代
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := []*TreeNode{root.Left, root.Right}
	for len(q) > 1 {
		if q[0] != nil || q[1] != nil {
			if (q[0] == nil && q[1] != nil) || (q[0] != nil && q[1] == nil) || q[0].Val != q[1].Val {
				return false
			}
			q = append(q, q[0].Left, q[1].Right, q[0].Right, q[1].Left)
		}
		q = q[2:]
	}
	return len(q) == 0
}

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
