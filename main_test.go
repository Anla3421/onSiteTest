package main

import (
	product "test/myproduct"
	"testing"
)

func BenchmarkMyProduct(b *testing.B) {
	input := []int{2, 1, 5, 9}
	input2 := []int{0, 1, 5, 9}
	input3 := []int{0, 0, 5, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		product.MyProduct(input)
	}
	for i := 0; i < b.N; i++ {
		product.MyProduct(input2)
	}
	for i := 0; i < b.N; i++ {
		product.MyProduct(input3)
	}
}

func BenchmarkMyProduct2(b *testing.B) {
	input := []int{2, 1, 5, 9}
	input2 := []int{0, 1, 5, 9}
	input3 := []int{0, 0, 5, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		product.MyProduct2(input)
	}
	for i := 0; i < b.N; i++ {
		product.MyProduct2(input2)
	}
	for i := 0; i < b.N; i++ {
		product.MyProduct2(input3)
	}
}

func BenchmarkMyProduct3(b *testing.B) {
	input := []int{2, 1, 5, 9}
	input2 := []int{0, 1, 5, 9}
	input3 := []int{0, 0, 5, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MyProduct3(input)
	}
	for i := 0; i < b.N; i++ {
		MyProduct3(input2)
	}
	for i := 0; i < b.N; i++ {
		MyProduct3(input3)
	}
}
