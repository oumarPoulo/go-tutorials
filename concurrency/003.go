package main

import (
	"fmt"
	"sync"

	"github.com/oumarPoulo/go-tutorials/structs"
)

// Waiting for multiple goroutines to finish without using Scanln
// If we need to wait for multiple goroutines to complete, sync.WaitGroup goes to the rescue.
// A WaitGroup allows to wait for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the worker goroutines runs and calls Done when finished.
// At the same time, Wait can be used to block until all goroutines have finished.
// The following example illustrates it.
func main() {
	var peoples = structs.Persons{
		{
			Age:      18,
			Name:     "Alpha",
			Nickname: "OumarPoulo",
		},
		{
			Age:      14,
			Name:     "Bambewel",
			Nickname: "Majesty",
		},
		{
			Age:      15,
			Name:     "Aissatou",
			Nickname: "Flash",
		},
		{
			Age:      10,
			Name:     "Idrissa",
			Nickname: "Paikoun",
		},
	}

	var wg sync.WaitGroup

	for i := 0; i < len(peoples); i++ {
		fmt.Println("current person == key:", i, " people=", peoples[i])
		wg.Add(1)
		go func(p structs.Person, s *sync.WaitGroup) {
			defer s.Done()
			p.Talk()
		}(peoples[i], &wg)
	}

	fmt.Println("Main: execute other functions before goroutines finish")
	wg.Wait()
	fmt.Println("Main: Waiting for goroutines to finish")
	fmt.Println("Main: Completed")
}

// current person == key: 0  people= {18 Alpha OumarPoulo}
// current person == key: 1  people= {14 Bambewel Majesty}
// current person == key: 2  people= {15 Aissatou Flash}
// current person == key: 3  people= {10 Idrissa Paikoun}
// Main: execute other functions before goroutines finish
// Hello my name is Idrissa, i'm 10 years old and my nickname is Paikoun booya
// Hello my name is Bambewel, i'm 14 years old and my nickname is Majesty booya
// Hello my name is Aissatou, i'm 15 years old and my nickname is Flash booya
// Hello my name is Alpha, i'm 18 years old and my nickname is OumarPoulo booya
// Main: Waiting for goroutines to finish
// Main: Completed
