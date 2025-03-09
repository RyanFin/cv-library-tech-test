package main

import (
	"cv-library-tech-test/pkg/model"
	"fmt"
	"time"
)

func initBeeHive() {

}

func main() {
	fmt.Println("hello world")

	qb := model.NewQueenBee()

	// create a new player instance
	player := model.NewPlayer()

	fmt.Println(qb)

	// Create a game loop
	var name string
	fmt.Print("enter your first name: ")
	fmt.Scan(&name)

	player.Name = name

	fmt.Println(player)

	// start the game loop
	start(player)

}

func start(player model.Player) {
	fmt.Println("Actions: Health [h], View Beehive[b], Attack[a]")

	for player.Health > 0 {
		var input string
		fmt.Print("Enter action: ")
		fmt.Scan(&input) // Reads a single word input

		switch input {
		case "h":
			fmt.Printf("%s's Health: %d\n", player.Name, player.Health)
		case "b":
			fmt.Println("Viewing the beehive...")
		case "a":
			fmt.Println("Attacking the beehive...")
		default:
			fmt.Println("Invalid action. Try again.")
		}

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Game Over! You have no health left.")

}
