package goroutine

import (
	"fmt"
	"sync"
)

// 并发爬虫10000

func PaMain() {
	var wait sync.WaitGroup
	var nums chan bool
	var strs chan bool

	wait.Add(2)
	i := 0
	// 数字
	for {
		select {
		case <-nums:
			fmt.Println("数字：", i)
			i++
			strs <- true
			wait.Done()
		}
	}

	// 字母
	for {
		z := 'A'
		select {
		case <-strs:
			if z > 'Z' {
				close(strs)
				return
			}
			fmt.Println("字母：", z)
			z++
			nums <- true
			wait.Done()
		}
	}
	// 字母
	nums <- true
	wait.Wait()
}
