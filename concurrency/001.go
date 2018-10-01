package main

import "fmt"

type Person struct {
	Age      int
	Name     string
	Nickname string
}

type Persons []Person

func (p Person) ToString() string {
	return fmt.Sprintf("Name=%s Nickname=%s, Age=%d", p.Name, p.Nickname, p.Age)
}

func (p *Person) Talk() {
	fmt.Printf("Hello my name is %s, i'm %d years old and my nickname is %s booya\n", p.Name, p.Age, p.Nickname)
}

var peoples = Persons{
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

func main() {
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
