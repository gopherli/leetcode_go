package leetcode

import (
	"log"
	"testing"
)

func BenchmarkLc2169(b *testing.B) {
	num1, num2 := 2, 3
	count := CountOperations(num1, num2)
	log.Println("count:", count)
}
