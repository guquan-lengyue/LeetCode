package main

// import "fmt"

// func min(arr ...int) int {
// 	minInt := 0x3f3f3f
// 	for i := 0; i < len(arr); i++ {
// 		if minInt > arr[i] {
// 			minInt = arr[i]
// 		}
// 	}
// 	return minInt
// }

// func coinChange(coins []int, amount int) int {
// 	// dp[i] 表示 面额为i是使用的零钱数
// 	dp := make([]int, amount+1)
// 	// 初始化值为最大值
// 	// 方便后续最优解处理
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = amount + 1
// 	}
// 	// 总金额为0时不需要零钱数量
// 	dp[0] = 0

// 	for i := 1; i <= amount; i++ {
// 		for j := 0; j < len(coins); j++ {

// 			if coins[j] <= i {
// 				// 如果硬币面额小于当前总金额
// 				// 则在取或不取中得到最优解
// 				dp[i] = min(dp[i], dp[i-coins[j]]+1)
// 			}
// 			// 如果当前零钱面额大于当前总金额
// 			// 不做任何处理
// 		}
// 	}
// 	// 如果当前总金额需要零钱数量没变化
// 	// 则认为当前总金额没有任何零钱组合
// 	if dp[amount] > amount {
// 		return -1
// 	}
// 	return dp[amount]
// }

// func main() {
// 	coins := []int{2}
// 	amount := 3
// 	rst := coinChange(coins, amount)
// 	fmt.Println(rst)
// }
