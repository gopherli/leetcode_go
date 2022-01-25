package go_t

import (
	"testing"
)

func LearnTest(t *testing.T) {
	// 子测试用例
	t.Run("pos", func(t *testing.T) {
		if Add(2, 3) != 5 {
			t.Fatal("fail")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Add(3, 3) == 6 {
			t.Errorf("fail")
		}
	})
}
