package main

import (
	"fmt"

	"github.com/oumarPoulo/go-tutorials/structs"
)

func main() {
	// This program consists of two goroutines.
	// The first goroutine is implicit and is the main function itself.
	// The second goroutine is created when we call go Talk method.
	// Normally when we invoke a function our program will execute all the statements in a function and
	// then return to the next line following the invocation.
	// With a goroutine we return immediately to the next line and don't wait for the function to complete.
	// This is why the call to the Scanln function has been included;
	// without it the program would exit before being given the opportunity to print all the numbers.
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

	for i := 0; i < len(peoples); i++ {
		fmt.Println("current person == key:", i, " people=", peoples[i])
		go peoples[i].Talk()
	}

	// without it the program would exit before being given the opportunity to print all the numbers.
	var input string
	fmt.Scanln(&input)
}
