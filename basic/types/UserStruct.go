package main

import(
	"fmt"
)

type User struct {
	Id int
	Name, Location string
}

type Player struct {
	User
	GameId int
}

func main() {
	p := Player{
		User{Id: 1, Name: "2", Location: "3"},
		90000,
	}
	fmt.Println(p)
	fmt.Println(p.Id)
	p.Id = 100
	fmt.Println(p.Id)
}
