package main


import (
	"fmt"
	"sync"

	"runtime"

	"go/types"
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



data :=  dataArguments{}
	newThread3(&wg, PrintNumber3,data)


	wg.Wait()
}

type dataArguments struct {
	fields []*types.Var
	tags   []string
}

func PrintNumber3(any ...interface{})  {
	for int := 0; int < 27; int++ {
		fmt.Printf("%d ", int)
	}
}

type typed3 func(ah ...interface{})
func newThread3(wg *sync.WaitGroup,procesor typed3,obj dataArguments)  {
	wg.Add(1)
	go func() {
		defer wg.Done()
		procesor()
	}()
}
