package main

import (
	"fmt"
	"strings"
)

type Renderer interface {
	RenderGrid()
}

type renderer struct {
	grid Grid
}

func NewRenderer(gridToRender Grid) Renderer {
	return &renderer{
		grid: gridToRender,
	}
}

func (r *renderer) RenderGrid() {
	for _, row := range r.grid.GetTileGrid() {
		fmt.Printf("%s+\n", strings.Repeat("+------", r.grid.GetSize()))
		for _, colItem := range row {
			if colItem != nil {
				fmt.Printf("|  %2v  ", colItem.GetId())
			} else {
				fmt.Printf("|  %2v  ", "%")
			}

		}
		fmt.Print("|\n")
	}
	fmt.Printf("%s+\n", strings.Repeat("+------", r.grid.GetSize()))
}
