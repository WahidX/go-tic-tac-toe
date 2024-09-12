package grid

var gridSize = 3 // default

type grid struct {
	mat   [][]string
	nextX bool
}

type Grid interface {
	Mark(pos *Position) (res TurnResult)
}

func NewGrid(size int) Grid {
	if size > 0 {
		gridSize = size
	}

	g := &grid{
		mat:   make([][]string, gridSize),
		nextX: true,
	}

	for i := range g.mat {
		g.mat[i] = make([]string, gridSize)
	}

	return g
}

func (g *grid) Mark(pos *Position) (res TurnResult) {
	if !pos.validate() {
		return InvalidMove
	}

	if g.mat[pos.X][pos.Y] != "" {
		return InvalidMove
	}

	if g.nextX {
		g.mat[pos.X][pos.Y] = "X"
	} else {
		g.mat[pos.X][pos.Y] = "O"
	}

	g.nextX = !g.nextX

	return g.checkMove(pos)
}

func (g *grid) checkMove(lastPos *Position) (res TurnResult) {
	// Horizontal
	for i := 0; i < gridSize; i++ {
		if g.mat[lastPos.X][i] != g.mat[lastPos.X][lastPos.Y] {
			return HorizontalWin
		}
	}

	// Vertical
	for i := 0; i < gridSize; i++ {
		if g.mat[i][lastPos.Y] != g.mat[lastPos.X][lastPos.Y] {
			return VerticalWin
		}
	}

	// Right Diagonal
	for i := 0; i < gridSize; i++ {
		if g.mat[i][i] != g.mat[lastPos.X][lastPos.Y] {
			return RDiagonalWin
		}
	}

	// Left Diagonal
	for i := 0; i < gridSize; i++ {
		if g.mat[i][gridSize-i-1] != g.mat[lastPos.X][lastPos.Y] {
			return LDiagonalWin
		}
	}

	return ValidMove
}
