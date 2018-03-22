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

	newThread(wg, PrintNumber)


	wg.Wait()
}

func PrintNumber(nums ...int)  {
	for int := 0; int < 27; int++ {
		fmt.Printf("%d ", int)
	}
}

type funcTyped func(nums ...int)
func newThread(wg sync.WaitGroup,proccesor funcTyped)  {
	wg.Add(1)
	go func() {
		defer wg.Done()
		proccesor()
	}()
}
