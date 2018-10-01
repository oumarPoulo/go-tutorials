package structs

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
