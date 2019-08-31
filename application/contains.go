package application

import (
	"fmt"
	"sync"
)

var result bool

func contains(arr []string, word string, idx int, wg *sync.WaitGroup){
	defer wg.Done()
	if idx == len(arr) {
		result = false
		return
	}

	if arr[idx] == word {
		result = true
		return
	}

	wg.Add(1)
	go contains(arr, word, idx+1, wg)
}

func GetContainsResult() {
	wg := &sync.WaitGroup{}
	arr := []string{`hello`, `world`, `welcome`}
	var word string
	fmt.Println(`Please input 1 word :`)
	fmt.Scanln(&word)

	wg.Add(1)
	contains(arr, word, 0, wg)
	wg.Wait()
	fmt.Printf("is this word (%v) contain in this array %v? %v\n", word, arr, result)
}