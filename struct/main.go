package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Username string `json:"username"`
	Level    uint   `json:"level"`
	Status   string `json:"status"`
	Class    string `json:"class"`
}

func (p *Player) LevelUp() {
	p.Level++
}

func (p Player) GetUsername() string {
	return p.Username
}

func main() {
	player1 := Player{
		Username: "player1",
		Level:    1,
		Status:   "active",
		Class:    "warrior",
	}

	fmt.Println(player1.GetUsername())

	player1.LevelUp()

	jsonData, _ := json.MarshalIndent(&player1, "", "\t")
	fmt.Println(string(jsonData))

	fmt.Println(player1)

	fmt.Println(player1.Level)
}
