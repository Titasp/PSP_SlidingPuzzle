package main

import (
	"fmt"
	"github.com/titasp/PSP_SlidingPuzzle/input"
	"github.com/titasp/PSP_SlidingPuzzle/layout"
	"github.com/titasp/PSP_SlidingPuzzle/rendering"
	"os"
	"strconv"
)

const (
	GridSize = 3
)

// Strategy pattern

type textRenderer struct {
}

func (r *textRenderer) Render() {
	fmt.Println("this is text renderer")
}

func main() {

	// Create tile layout
	grid, err := layout.NewGrid(GridSize)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ======= Singleton pattern =========
	renderer := rendering.NewRenderer(grid)
	renderer = rendering.NewRenderer(nil)
	// ===================================

	// ======= Strategy pattern (behavioural pattern) =======
	//renderer = rendering.Renderer(&textRenderer{})

	// Generate valid id for validation based on configured layout size
	var validIds []string
	for i := 1; i <= GridSize*GridSize; i++ {
		validIds = append(validIds, strconv.FormatInt(int64(i), 10))
	}

	// Create input handlers with validation
	tileIdInputHandler := input.NewHandler(
		"Please enter tile ID: ",
		validIds...,
	)
	movementDirectionInputHandler := input.NewHandler("Please enter movement direction ('u', 'd', 'l', 'r'): ",
		"u",
		"d",
		"l",
		"r")

	for {
		//// Clear console
		//cmd := exec.Command("cmd", "/c", "cls")
		//cmd.Stdout = os.Stdout
		//cmd.Run()

		renderer.Render()

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
		finished, err := grid.MoveTile(idCommand, layout.MoveDirection(moveCommand))
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
