package main

import (
	"fmt"
	"time"
)

type Interface1 []string
type Interface2 []string

func processRandom(data interface{}, id int) {
	for i := 0; i < 4; i++ {
		go fmt.Printf("%v %d\n", data, i+1)
	}
}

func main() {
	data1 := Interface1{"coba1", "coba2", "coba3"}
	data2 := Interface2{"bisa1", "bisa2", "bisa3"}

	processRandom(data1, 1)
	processRandom(data2, 2)

	time.Sleep(2 * time.Second)
}
