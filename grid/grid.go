package grid

var gridSize = 3 // default

type grid struct {
	mat   [][]string
	nextX bool
}

type Grid interface {
	Mark(pos *Position) (res TurnResult)
	Size() int
	GetMat() [][]string
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

func (g *grid) GetMat() [][]string {
	return g.mat
}

func (g *grid) Size() int {
	return gridSize
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
	var (
		horizontal = true
		vertical   = true
		rDiagonal  = true
		lDiagonal  = true
	)

	// Horizontal
	for i := 0; i < gridSize; i++ {
		if g.mat[lastPos.X][i] != g.mat[lastPos.X][lastPos.Y] {
			horizontal = false
			break
		}
	}
	if horizontal {
		return HorizontalWin
	}

	// Vertical
	for i := 0; i < gridSize; i++ {
		if g.mat[i][lastPos.Y] != g.mat[lastPos.X][lastPos.Y] {
			vertical = false
			break
		}
	}
	if vertical {
		return VerticalWin
	}

	// Left Diagonal
	for i := 0; i < gridSize; i++ {
		if g.mat[i][i] != g.mat[lastPos.X][lastPos.Y] {
			lDiagonal = false
			break
		}
	}
	if lDiagonal {
		return LDiagonalWin
	}

	// Right Diagonal
	for i := 0; i < gridSize; i++ {
		if g.mat[i][gridSize-i-1] != g.mat[lastPos.X][lastPos.Y] {
			rDiagonal = false
			break
		}
	}
	if rDiagonal {
		return RDiagonalWin
	}

	return ValidMove
}
