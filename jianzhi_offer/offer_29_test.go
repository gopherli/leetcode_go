package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer29(t *testing.T) {
	twoArray := [][]int{
		{1},
		{10},
		{9},
	}

	orderArray := SpiralOrder(twoArray)
	log.Printf("顺时针打印矩阵：%v", orderArray)
}
