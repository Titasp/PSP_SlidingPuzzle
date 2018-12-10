package layout

type Tile interface {
	GetPosition() (int, int)
	SetPosition(int, int)
	GetId() string
}

type tile struct {
	id   string
	posX int
	posY int
}

func NewTile(id string) Tile {
	return &tile{
		id:   id,
		posX: -1,
		posY: -1,
	}
}

func (t *tile) GetPosition() (int, int) {
	return t.posX, t.posY
}

func (t *tile) SetPosition(posX, posY int) {
	t.posX = posX
	t.posY = posY
}

func (t *tile) GetId() string {
	return t.id
}
