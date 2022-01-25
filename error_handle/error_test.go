package error_handle

import (
	"fmt"
	"testing"
)

func TestError_leran(t *testing.T) {
	stringOne := "my name:.lizhi."
	fmt.Println(stringOne)

	stringTwo := `my name:"lizhi"`
	fmt.Println(stringTwo)

	const msg = "unsnapact: %v, %v"
	fmt.Printf(msg, 1, 2)
}
