package layout

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

var (
	ErrInvalidDirection              = errors.New("invalid direction")
	ErrInvalidTileId                 = errors.New("invalid tile id")
	ErrInvalidGridSize               = errors.New("invalid layout size, must be larger than 1")
	ErrInvalidMoveAction_OutOfBounds = errors.New("invalid move action, this would move tile out of layout")
	ErrInvalidMoveAction_TileBlocked = errors.New("invalid move action, this tile is blocked")
)

type MoveDirection string

const (
	MoveUp    = MoveDirection("u")
	MoveDown  = MoveDirection("d")
	MoveLeft  = MoveDirection("l")
	MoveRight = MoveDirection("r")
)

type Grid interface {
	MoveTile(string, MoveDirection) (bool, error)
	GetSize() int
	GetTileGrid() [][]Tile
}

type grid struct {
	size     int
	tiles    map[string]Tile
	tileGrid [][]Tile
}

func NewGrid(size int) (Grid, error) {
	if size < 2 {
		return nil, ErrInvalidGridSize
	}

	gridInstance := grid{
		size:     size,
		tileGrid: make([][]Tile, size),
		tiles:    make(map[string]Tile),
	}

	for i := range gridInstance.tileGrid {
		gridInstance.tileGrid[i] = make([]Tile, size)
	}

	tileIds := make([]int, size*size-1)
	for i := 1; i < size*size; i++ {
		tileIds[i-1] = i
	}

	// Shuffle tile ids
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(tileIds), func(i, j int) {
		tileIds[i], tileIds[j] = tileIds[j], tileIds[i]
	})

	row := 0
	col := 0

	// Create tiles and populate layout
	for i := 0; i < size*size; i++ {
		if i == size*size-1 {
			gridInstance.tileGrid[row][col] = nil
			break
		}
		tile := NewTile(strconv.FormatInt(int64(tileIds[i]), 10))
		tile.SetPosition(col, row)
		gridInstance.tiles[tile.GetId()] = tile
		gridInstance.tileGrid[row][col] = tile

		if col/(size-1) == 1 {
			col = 0
			row++
		} else {
			col++
		}
	}

	return &gridInstance, nil
}

func (g *grid) MoveTile(tileId string, direction MoveDirection) (bool, error) {

	tileToMove, exists := g.tiles[tileId]
	if !exists {
		return false, ErrInvalidTileId
	}

	originalPosX, originalPosY := tileToMove.GetPosition()
	posX, posY := originalPosX, originalPosY

	switch direction {
	case "u":
		posY -= 1
		break
	case "d":
		posY += 1
		break
	case "l":
		posX -= 1
		break
	case "r":
		posX += 1
		break
	default:
		return false, ErrInvalidDirection
	}

	if !(posX >= 0 && posX < g.size && posY >= 0 && posY < g.size) {
		return false, ErrInvalidMoveAction_OutOfBounds
	}

	tileAtNewPosition := g.tileGrid[posY][posX]
	if tileAtNewPosition != nil {
		return false, ErrInvalidMoveAction_TileBlocked
	}

	tileToMove.SetPosition(posX, posY)
	g.tileGrid[posY][posX] = tileToMove
	g.tileGrid[originalPosY][originalPosX] = nil

	return g.isFinished(), nil
}

func (g *grid) GetSize() int {
	return g.size
}

func (g *grid) GetTileGrid() [][]Tile {
	return g.tileGrid
}

func (g *grid) isFinished() bool {
	lastNum := 0
	for row := 0; row < g.size; row++ {
		for col := 0; col < g.size; col++ {
			tile := g.tileGrid[row][col]
			if tile == nil && (row != g.size-1 || col != g.size-1) {
				return false
			}

			if row != g.size-1 && col != g.size-1 {
				tileId, _ := strconv.ParseInt(tile.GetId(), 10, 64)
				if lastNum > int(tileId) {
					return false
				}
				if tile != nil {
					lastNum = int(tileId)
				}
			}

		}
	}
	return true
}
