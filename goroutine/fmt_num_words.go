package goroutine

import (
	"fmt"
)

// 交替打印数字和字母

func FmtNumWords() {
	letter, number := make(chan bool), make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()
	// wait.Add(1)
	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					close(letter)
					// close(number)
					return
				}
				fmt.Print(string(i))
				i++
				number <- true
			}

		}
	}()
	number <- true
}
