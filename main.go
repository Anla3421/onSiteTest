package main

import (
	"fmt"
)

func main() {
	// process.Process([]int{2, 1, 5, 9}, 2)

	// 測試 MyProduct 函數
	result := MyProduct([]int{2, 1, 5, 9})
	fmt.Println(result) // 輸出：[45 90 18 10]
}

// MyProduct 返回一個新的切片，其中每個元素是其他所有元素的乘積
func MyProduct(nums []int) []int {
	n := len(nums) // 4
	if n == 0 {
		return []int{}
	}

	// 創建結果切片
	result := make([]int, n)

	// 第一次遍歷：計算左側所有數的乘積
	leftProduct := 1
	for i := 0; i < n; i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
		fmt.Println("result", result)
		fmt.Println("leftProduct", leftProduct)
	}

	// 第二次遍歷：計算右側所有數的乘積並與左側乘積相乘
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
		fmt.Println("result", result)
		fmt.Println("rightProduct", rightProduct)
	}

	return result
}
