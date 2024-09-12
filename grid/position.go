package grid

type Position struct {
	X int
	Y int
}

func (p *Position) validate() bool {
	return p.X >= 0 && p.X < gridSize && p.Y >= 0 && p.Y < gridSize
}
