package main

import (
	"fmt"

	"github.com/oumarPoulo/go-tutorials/structs"
)

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
	// peoples will talk sequentially
	for key, people := range peoples {
		fmt.Println("current person == key:", key, " people=", people)
		people.Talk()
	}
	// current person == key: 0  people= {18 Alpha OumarPoulo}
	// Hello my name is Alpha, i'm 18 years old and my nickname is OumarPoulo booya
	// current person == key: 1  people= {14 Bambewel Majesty}
	// Hello my name is Bambewel, i'm 14 years old and my nickname is Majesty booya
	// current person == key: 2  people= {15 Aissatou Flash}
	// Hello my name is Aissatou, i'm 15 years old and my nickname is Flash booya
	// current person == key: 3  people= {10 Idrissa Paikoun}
	// Hello my name is Idrissa, i'm 10 years old and my nickname is Paikoun booya
}
