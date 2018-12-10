package rendering

import (
	"fmt"
	"github.com/titasp/PSP_SlidingPuzzle/layout"
	"strings"
	"sync"
)

var (
	once     sync.Once
	instance Renderer
)

type Renderer interface {
	Render()
}

type gridRenderer struct {
	grid layout.Grid
}

func NewRenderer(gridToRender layout.Grid) Renderer {
	once.Do(func() {
		instance = &gridRenderer{grid: gridToRender}
	})
	return instance
}

func (r *gridRenderer) Render() {
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
