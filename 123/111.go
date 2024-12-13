package product

// Implement a function F that takes a slice of integers (A).
// It then returns slice (B) where B[i] is the product of all A[j] where j != i.
// For example: If A = {2, 1, 5, 9}, then B would be {45, 90, 18, 10}.
// Please implement it in the time complexity of O(n).

// 獲取長度的加總的數量（如1+2+3+4+...)
// 建立一個新slice用以儲存位置跟總數的差值（c:=[9,8,7,6])
// 如果b現在要計算的index+1跟c[index]它加起來等於10，那在計算數字的時候就當作1，b[0]=1*a[1]*a[2]*a[3],  b[1]=a[0]*1*a[2]*a[3])
// 想辦法不要O(n)^2
