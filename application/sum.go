package application

import (
	"fmt"
	"sync"
)

var total int64 = 0

func sum(arr []int64, idx int, wg *sync.WaitGroup){
	defer wg.Done()

	if idx == len(arr) {
		return
	}
	total += arr[idx]

	wg.Add(1)
	go sum(arr, idx+1, wg)
}

func GetSumResult() {
	wg := &sync.WaitGroup{}
	arr := []int64{1,2,3,4,5}

	wg.Add(1)
	sum(arr, 0, wg)
	wg.Wait()
	fmt.Printf("Summary of array %v is %d\n", arr, total)
}