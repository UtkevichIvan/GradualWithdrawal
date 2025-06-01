package main

import (
	"fmt"
	"sync"
	"time"
)

type (
	WriteText func(string)
	Wait      func()
)

func main() {
	writeText, wait := TypeWriter(time.Millisecond * 200)
	defer wait()
	writeText("The matrix has you\n")
	writeText("Per aspera ad astra\n")

	fmt.Print("\n[Проверка что ниче не блокируется]\n")

}

func TypeWriter(delay time.Duration) (WriteText, Wait) {
	ch := make(chan string, 100)
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			val, ok := <-ch
			if !ok {
				break
			}
			fmt.Print(val)
			time.Sleep(delay)
		}
	}()
	return func(s string) {
			for _, char := range s {
				ch <- string(char)
			}
		}, func() {
			close(ch)
			wg.Wait()
		}
}
