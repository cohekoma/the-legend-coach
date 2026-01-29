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

func (team *Team) PrintSquad() {
	fmt.Printf("--- Squad List of %s ---\n", team.Name)

	for index, player := range team.Players {
		fmt.Printf("[%d] Player info: %s - %s - %d\n", index, player.Name, player.Position, player.Overall)
	}
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

	t1.PrintSquad()
}
