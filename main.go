package main

import (
	"flag"
	"fmt"
	"github.com/titasp/PSP_SlidingPuzzle/input"
	"github.com/titasp/PSP_SlidingPuzzle/layout"
	"github.com/titasp/PSP_SlidingPuzzle/rendering"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	GridSize = 3
)

var (
	rendererType int
)

// 'otherRenderer' and its usage is for strategy pattern showcase
type otherRenderer struct {
	grid layout.Grid
}

func (r *otherRenderer) Render() {
	for _, row := range r.grid.GetTileGrid() {
		fmt.Printf("%s+\n", strings.Repeat("+======", r.grid.GetSize()))
		for _, colItem := range row {
			if colItem != nil {
				fmt.Printf("|  %2v  ", colItem.GetId())
			} else {
				fmt.Printf("|  %2v  ", "%")
			}

		}
		fmt.Print("|\n")
	}
	fmt.Printf("%s+\n", strings.Repeat("+======", r.grid.GetSize()))
}

func init() {
	flag.IntVar(&rendererType, "renderer_type", 0, "enter renderer type")
	flag.Parse()

	if rendererType != 0 && rendererType != 1 {
		flag.PrintDefaults()
		log.Fatal("bad renderr type entered")
	}
}

func main() {

	// Create tile layout
	grid, err := layout.NewGrid(GridSize)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var renderer rendering.Renderer
	if rendererType == 0 {
		renderer = rendering.NewRenderer(grid)

		// ======= Singleton pattern =========
		// uncommenting of one line bellow shows that singleton pattern is working
		// and new call for 'NewRenderer' just returns the same instance
		//renderer = rendering.NewRenderer(nil)
	} else {
		renderer = rendering.Renderer(&otherRenderer{grid: grid})
	}

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
