package process

import (
	"fmt"
	"sync"
	"time"
)

// 以下共兩題
// 請實現一個函數，限制最多同時執行 n 個 goroutine
// 1. 需要限制同時執行的 goroutine 數量為 n
// 2. 每個 task 需要在獨立的 goroutine 中執行
// 3. 需要等待所有 task 完成

func Process(tasks []int, n int) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, n)

	for _, task := range tasks {
		wg.Add(1)
		go func(task int) {
			defer wg.Done()
			ch <- struct{}{}
			defer func() {
				<-ch
			}()
			processTask(task) // 執行任務指定任務.....
		}(task)
	}

	wg.Wait()
}

func processTask(task int) {
	fmt.Printf("task: %d\n", task)
	time.Sleep(3 * time.Second)
}
