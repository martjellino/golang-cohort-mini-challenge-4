package main

import (
	"fmt"
	"sync"
)

type Interface1 []string
type Interface2 []string

func main() {
	var wg sync.WaitGroup
	m := &sync.Mutex{}

	data1 := Interface1{"coba1", "coba2", "coba3"}
	data2 := Interface2{"bisa1", "bisa2", "bisa3"}

	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			m.Lock()
			fmt.Println(data1, i+1)
			fmt.Println(data2, i+1)
			m.Unlock()
		}
	}()

	wg.Wait()
}
