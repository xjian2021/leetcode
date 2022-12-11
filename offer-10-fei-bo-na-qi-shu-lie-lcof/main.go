package main

import "fmt"

//剑指 Offer 10- I. 斐波那契数列
/*
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：

F(0) = 0,   F(1) = 1, F(2) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：
输入：n = 2
输出：1

示例 2：
输入：n = 5
输出：5
*/

func main() {
	n := 95
	fmt.Println(fib2(n))
}

// fib 递归(LeetCode会超时)
func fib(n int) int {
	if n < 2 {
		return n
	}
	return (fib(n-1) + fib(n-2)) % 1000000007
}

// fib2 DP动态规划
func fib2(n int) int {
	if n < 2 {
		return n
	}
	// 利用递推公式:f(n) = f(n-1)+f(n-2)进行模拟计算过程
	o, r := 0, 1
	for i := 2; i <= n; i++ {
		o, r = r, (r+o)%1000000007
	}
	return r
}

// TODO 还有一种线性代数式的解题，没搞懂...
