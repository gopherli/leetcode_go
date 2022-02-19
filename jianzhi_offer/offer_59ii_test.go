package jianzhi_offer

import (
	"fmt"
	"testing"
)

func BenchmarkOffer59ii(b *testing.B) {
	this := Constructor()
	fmt.Println("Max_value:", this.Max_value())
	fmt.Println("Pop_front:", this.Pop_front())
	fmt.Println("Push_back:", 1)
	this.Push_back(1)
	// fmt.Println("Push_back:", 2)
	// this.Push_back(2)
	// fmt.Println("Push_back:", 1)
	// this.Push_back(1)
	fmt.Println("Max_value:", this.Max_value())
	fmt.Println("Max_value:", this.Max_value())
	fmt.Println("Pop_front:", this.Pop_front())
	fmt.Println("Max_value:", this.Max_value())
}
