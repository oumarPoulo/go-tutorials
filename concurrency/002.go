package main

import (
	"fmt"

	"github.com/oumarPoulo/go-tutorials/structs"
)

// This program consists of two goroutines.
// The first goroutine is implicit and is the main function itself.
// The second goroutine is created when we call go Talk method.
// Normally when we invoke a function our program will execute all the statements in a function and
// then return to the next line following the invocation.
// With a goroutine we return immediately to the next line and don't wait for the function to complete.
// This is why the call to the Scanln function has been included;
// without it the program would exit before being given the opportunity to print all the numbers.
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

	for i := 0; i < len(peoples); i++ {
		fmt.Println("current person == key:", i, " people=", peoples[i])
		go peoples[i].Talk()
	}

	// without it the program would exit before being given the opportunity to print all the numbers.
	var input string
	fmt.Scanln(&input)
}

// example of result
// current person == key: 0  people= {18 Alpha OumarPoulo}
// current person == key: 1  people= {14 Bambewel Majesty}
// current person == key: 2  people= {15 Aissatou Flash}
// current person == key: 3  people= {10 Idrissa Paikoun}
// Hello my name is Alpha, i'm 18 years old and my nickname is OumarPoulo booya
// Hello my name is Idrissa, i'm 10 years old and my nickname is Paikoun booya
// Hello my name is Bambewel, i'm 14 years old and my nickname is Majesty booya
// Hello my name is Aissatou, i'm 15 years old and my nickname is Flash booya
