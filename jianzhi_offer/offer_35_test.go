package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer35(t *testing.T) {
	head := &Node{
		Val:    1,
		Next:   &Node{Val: 2, Next: &Node{Val: 3, Next: nil, Random: nil}, Random: nil},
		Random: nil,
	}
	newHead := DeepCopy(head)
	log.Printf("复杂链表的深拷贝：%v", newHead)
}
