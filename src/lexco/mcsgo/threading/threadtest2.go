package main


import (
"fmt"
"sync"

"runtime"

)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()
		//time.Sleep (2 * time.Microsecond)
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)

		}
	}()

	go func() {
		defer wg.Done()
		//time.Sleep (2 * time.Microsecond)
		for number := 0; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()



	wg.Wait()
}
