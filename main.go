package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const (
	GridSize = 3
)

func main() {

	// Create tile grid
	grid, err := NewGrid(GridSize)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	renderer := NewRenderer(grid)

	// Generate valid id for validation based on configured grid size
	validIds := []string{}
	for i := 1; i <= GridSize*GridSize; i++ {
		validIds = append(validIds, strconv.FormatInt(int64(i), 10))
	}

	// Create input handlers with validation
	tileIdInputHandler := NewInputHandler(
		"Please enter tile ID: ",
		validIds...
	)
	movementDirectionInputHandler := NewInputHandler("Please enter movement direction ('u', 'd', 'l', 'r'): ",
		"u",
		"d",
		"l",
		"r")

	for {
		// Clear console
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		renderer.RenderGrid()

		// Get tile id and movement direction from console
		idCommand, err := tileIdInputHandler.GetCommand()
		if err != nil {
			fmt.Println("Please enter info correctly, error: ", err)
			continue
		}

		moveCommand, err := movementDirectionInputHandler.GetCommand()
		if err != nil {
			fmt.Println("Please enter info correctly, error: ", err)
			continue
		}

		// Move tile and check if tiles are arranged correctly
		finished, err := grid.MoveTile(idCommand, MoveDirection(moveCommand))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// If arranged correctly then show success text
		if finished {
			fmt.Println("Success. You finished the game!")
			os.Exit(1)
		}

	}

}
