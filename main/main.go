package main

import "fmt"

type Player struct {
	Name     string
	Position string
	Overall  int
}

type Team struct {
	Name    string
	Players []Player
}

func main() {
	p1 := Player{
		Name:     "Bruno Fernandes",
		Position: "CAM",
		Overall:  88,
	}

	t1 := Team{
		Name:    "Manchester United",
		Players: []Player{p1},
	}

	fmt.Printf("This is %s\n", t1.Name)
	fmt.Printf("The captain of the team is %s, he's playing at the role %s and he's rated with %d in overall!\n", p1.Name, p1.Position, p1.Overall)
}
