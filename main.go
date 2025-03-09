package main

import (
	"cv-library-tech-test/pkg/model"
	"fmt"
	"os"
	"time"
)

func initBeeHive() (model.QueenBee, map[string][]model.WorkerBee, map[string][]model.DroneBee) {
	const numOfWorkerBees = 5
	const numOfDroneBees = 25

	workerBeeMap := make(map[string][]model.WorkerBee)
	droneBeeMap := make(map[string][]model.DroneBee)

	workerBees := []model.WorkerBee{}
	droneBees := []model.DroneBee{}

	// generate worker bees
	for i := 0; i < numOfWorkerBees; i++ {
		wb := model.NewWorkerBee()
		beeName := fmt.Sprintf("wb%v", i)
		workerBees = append(workerBees, wb)
		workerBeeMap[beeName] = workerBees
	}

	// generate drone bees
	for i := 0; i < numOfDroneBees; i++ {
		db := model.NewDroneBee()
		beeName := fmt.Sprintf("db%v", i)
		droneBees = append(droneBees, db)
		droneBeeMap[beeName] = droneBees

	}

	return model.NewQueenBee(), workerBeeMap, droneBeeMap
}

func beehiveStatus() {

}

func main() {
	// create a new player instance
	player := model.NewPlayer()

	// Create a game loop
	var name string
	fmt.Print("enter your first name: ")
	fmt.Scan(&name)

	player.Name = name

	fmt.Println(player)

	// initialise the beehive
	queenBee, workerBees, droneBees := initBeeHive()

	// start the game loop
	start(player, queenBee, workerBees, droneBees)

}

func start(player model.Player, queenBee model.QueenBee, workerBees map[string][]model.WorkerBee, droneBees map[string][]model.DroneBee) {
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
			fmt.Printf("Queen Bee: %v\n", queenBee)
			fmt.Printf("Worker Bees: %v\n", len(workerBees))
			fmt.Printf("Drone Bees: %v\n", len(droneBees))
		case "a":
			fmt.Println("Attacking the beehive...")
			// attack function. Hit or miss
			// AttackHitOrMiss()
		default:
			fmt.Println("Invalid action. Try again.")
		}

		// add a delay
		time.Sleep(500 * time.Millisecond)

		// check if player has completed the game
		if queenBee.Health == 0 && len(workerBees) == 0 && len(droneBees) == 0 {
			fmt.Printf("Congratulations. You have defeated ALL bees in the hive\n")

			// exit program
			os.Exit(0)
		}
	}

	// if player dies
	fmt.Println("Game Over! You have no health left.")

}
