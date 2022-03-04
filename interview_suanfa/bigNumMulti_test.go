package interview_suanfa

import "testing"

func BenchmarkBigNumMulti(b *testing.B) {
	BigNumMulti("1", "2")
}
