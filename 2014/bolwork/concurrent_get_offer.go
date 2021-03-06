package main

import (
	"fmt"
	"strconv"
	"sync"
)

//START OMIT
func fetchOffer(id int) string {
	return strconv.Itoa(id)
}

func main() {
	ids := []int{9000000065433456, 800000005435343, 7000000011602483, 600000005432543}

	wg := new(sync.WaitGroup) // of fetching goroutines

	for index, each := range ids {
		wg.Add(1) // to the waitgroup

		go func(i int, id int) {
			fmt.Println(i, fetchOffer(each)) // do you see the BUG?
			wg.Done()
		}(index, each)
	}

	wg.Wait() // for all started goroutines to finish
}

//END OMIT
