package concurrency

import "fmt"

type Person struct {
	Age int
	Name string
	Nickname string
}

type Persons []Person

var developers = []Person{
	{
		Age:20,
		Name: "alpha",
		Nickname: "Ghost",
	},
}

func (p *Person) ToString() string  {

}

func (p *Person) Talk() {
	fmt.Printf("Hello my name is %s, i'm %d years old and my nickname is %s booya\n", p.Name, p.Age, p.Nickname)
}

func SayHello(persons Persons)  {
	for i := 0; i < len(persons); i++ {

	}
}
