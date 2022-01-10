package jianzhi_offer

import (
	"testing"
)

func TestOffer09(t *testing.T) {
	cq := Constructor09()
	cq.DeleteHead()
	cq.in.Push(2)
	cq.in.Push(3)
	cq.DeleteHead()
	cq.DeleteHead()
}
