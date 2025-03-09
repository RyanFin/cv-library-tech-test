package main

import (
	"cv-library-tech-test/pkg/model"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func initBeeHive() (model.QueenBee, []model.WorkerBee, []model.DroneBee) {
	const numOfWorkerBees = 5
	const numOfDroneBees = 25

	workerBees := []model.WorkerBee{}
	droneBees := []model.DroneBee{}

	// generate worker bees
	for i := 0; i < numOfWorkerBees; i++ {
		wb := model.NewWorkerBee()
		// beeName := fmt.Sprintf("wb%v", i)
		workerBees = append(workerBees, wb)
		// workerBeeMap[beeName] = workerBees
	}

	// generate drone bees
	for i := 0; i < numOfDroneBees; i++ {
		db := model.NewDroneBee()
		// beeName := fmt.Sprintf("db%v", i)
		droneBees = append(droneBees, db)
		// droneBeeMap[beeName] = droneBees

	}

	return model.NewQueenBee(), workerBees, droneBees
}

func beehiveStatus(queenBee model.QueenBee, workerBees []model.WorkerBee, droneBees []model.DroneBee) {
	fmt.Println("Viewing the beehive...")
	fmt.Printf("Worker Bees: %v\n", len(workerBees))
	fmt.Printf("Drone Bees: %v\n", len(droneBees))

	fmt.Printf("Queen Bee: %v\n", queenBee)
	fmt.Println("Worker Bees: ", workerBees)
	fmt.Println("Drone Bees: ", droneBees)
}

func main() {
	// create a new player instance
	player := model.NewPlayer()

	fmt.Println(
		`
__________                ________                       
\______   \ ____   ____  /  _____/_____    _____   ____  
 |    |  _// __ \_/ __ \/   \  ___\__  \  /     \_/ __ \ 
 |    |   \  ___/\  ___/\    \_\  \/ __ \|  Y Y  \  ___/ 
 |______  /\___  >\___  >\______  (____  /__|_|  /\___  >
        \/     \/     \/        \/     \/      \/     \/ `)

	// Create a game loop
	var name string
	fmt.Print("enter your first name: ")
	fmt.Scan(&name)

	player.Name = name

	// initialise the beehive
	queenBee, workerBees, droneBees := initBeeHive()

	// start the game loop
	start(player, queenBee, workerBees, droneBees)

}

func start(player model.Player, queenBee model.QueenBee, workerBees []model.WorkerBee, droneBees []model.DroneBee) {
	for player.Health > 0 {
		var input string
		fmt.Printf("Actions: %v's health (%v%%), View Beehive[v], Attack[a]\nEnter action: ", player.Name, player.Health)
		fmt.Scan(&input) // Reads a single word input

		switch input {
		case "h":
			fmt.Printf("%s's Health: %d\n", player.Name, player.Health)
		case "v":
			beehiveStatus(queenBee, workerBees, droneBees)
		case "a":
			fmt.Println("Attacking the beehive...")
			AttackHitOrMiss(&player, &queenBee, &workerBees, &droneBees)
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

func AttackHitOrMiss(player *model.Player, queenBee *model.QueenBee, workerBees *[]model.WorkerBee, droneBees *[]model.DroneBee) {
	// +1 is the queen bee
	totalNumOfBees := len(*workerBees) + len(*droneBees) + 1

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Randomly select a bee
	randomIndex := rand.Intn(totalNumOfBees)

	// Pointer to the chosen bee
	var selectedBee *model.Bee

	if randomIndex == 0 {
		selectedBee = &queenBee.Bee
	} else if randomIndex <= len(*workerBees) {
		selectedBee = &(*workerBees)[randomIndex-1].Bee
	} else {
		selectedBee = &(*droneBees)[randomIndex-len(*workerBees)-1].Bee
	}

	// Randomly decide the outcome: player attacks, bee attacks, or miss
	action := rand.Intn(4) // 0: Player Attacks, 1: Bee Attacks, 2: Miss

	switch action {
	case 0:
		// Player attacks
		selectedBee.HitByPlayer()
		fmt.Printf("Player attacked the %s Bee!\n", selectedBee.BeeType)
		if selectedBee.Health <= 0 {
			// Remove the bee from its respective array
			fmt.Printf("The %s Bee has been defeated!\n", selectedBee.BeeType)
			switch selectedBee.BeeType {
			case "Queen":
				// Queen bee does not get removed, keep it in the hive
			case "Worker":
				removeWorkerBee(workerBees, randomIndex-1)
			case "Drone":
				removeDroneBee(droneBees, randomIndex-len(*workerBees)-1)
			}
		}

	case 1:
		// Bee attacks player
		selectedBee.HitsPlayer(player)
		fmt.Printf("The %s Bee attacked the player!\n", selectedBee.BeeType)
	case 2:
		// Miss
		fmt.Println("Player attack missed! No damage dealt.")
	case 3:
		fmt.Println("Bee attack missed the player! No damage received.")
	}

	// Print the bee's health after the action
	fmt.Printf("%s Bee's remaining health: %d\n", selectedBee.BeeType, selectedBee.Health)
	fmt.Printf("Player's remaining health: %d\n", player.Health)
}

// Helper function to remove a worker bee from the slice by index
func removeWorkerBee(workerBees *[]model.WorkerBee, index int) {
	*workerBees = append((*workerBees)[:index], (*workerBees)[index+1:]...)
}

// Helper function to remove a drone bee from the slice by index
func removeDroneBee(droneBees *[]model.DroneBee, index int) {
	*droneBees = append((*droneBees)[:index], (*droneBees)[index+1:]...)
}
