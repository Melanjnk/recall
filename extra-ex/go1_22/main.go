package main

import (
	"fmt"
	"time"
)

// что Выведет на экран этот код?
// что поменялось начиная с до 1.22?
func main() {

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 5)
}

//output in 1.22 is random order
//1
//3
//2
//4
//0
