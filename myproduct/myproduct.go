package product

import "fmt"

// Implement a function F that takes a slice of integers (A).
// It then returns slice (B) where B[i] is the product of all A[j] where j != i.
// For example: If A = {2, 1, 5, 9}, then B would be {45, 90, 18, 10}.
// Please implement it in the time complexity of O(n).

// 想辦法不要O(n)^2

// a := [2, 1, 5, 9]
// b := [45, 90, 18, 10]

// b[0]=a[1]*a[2]*a[3]
// len(a)=4

// 建立一個新slice用以儲存位置跟總數的差值（refSlice:=[9,8,7,6])
// 如果b現在要計算的index+1跟c[index]它加起來等於10(total)，
// 那在計算數字的時候就當作1，即 b[0]=1*a[1]*a[2]*a[3], b[1]=a[0]*1*a[2]*a[3])
// 可以處理input帶有0的slice，雖可以一套func做完(即不需要額外寫遇到0要怎麼辦)，但實作邏輯上比較複雜
// 空間複雜度應有符合要求
func MyProduct(nums []int) []int {
	lengthOfInput := len(nums)
	refSlice := make([]int, lengthOfInput)

	// 獲取input長度的加總的數量（如1+2+3+4+...)，用三角形面積公式計算
	total := lengthOfInput * (lengthOfInput + 1) / 2
	// 建立一個新slice用以儲存位置跟總數的差值（refSlice:=[9,8,7,6])
	for i := 1; i <= lengthOfInput; i++ {
		refSlice[i-1] = total - i
	}

	b := make([]int, lengthOfInput)
	for i := 0; i < lengthOfInput; i++ {
		needToTime := 1
		for index, v := range nums {
			// 如果b現在要計算的index+1跟c[index]它加起來等於10(total)，
			// 那在計算數字的時候就當作1，即 b[0]=1*a[1]*a[2]*a[3], b[1]=a[0]*1*a[2]*a[3])
			if index+1+refSlice[i] != total {
				needToTime *= v
			}
		}
		b[i] = needToTime
	}
	fmt.Println("input nums: ", nums)
	fmt.Println("result b: ", b)
	return b
}

// 把input slince全部乘起來再除掉要填入的該index的數字後變成b[index]即完成
// 雖然實作邏輯上比較單純，但需要額外處理slice內有0的情況
func MyProduct2(nums []int) []int {
	lengthOfInput := len(nums)
	b := make([]int, lengthOfInput)
	total := 1
	zeroCount := []int{}
	for k, v := range nums {
		if v == 0 {
			zeroCount = append(zeroCount, k)
			continue
		}
		total *= v
	}

	// 針對input slice有不同的0的數量來處理對應的情況
	switch len(zeroCount) {
	case 0:
		for i := 0; i < lengthOfInput; i++ {
			b[i] = total / nums[i]
		}
	// 有一個0，除了該index的value以外，都是0
	case 1:
		zeroIndex := zeroCount[0]
		for i := 0; i < lengthOfInput; i++ {
			if i == zeroIndex {
				b[i] = total
				continue
			}
			b[i] = 0
		}
	// 有兩個以上的0，則都是0
	default:
	}
	return b
}
