package grid

type TurnResult string

const (
	HorizontalWin TurnResult = "Horizontal Win"
	VerticalWin   TurnResult = "Vertical Win"
	RDiagonalWin  TurnResult = "Right Diagonal Win"
	LDiagonalWin  TurnResult = "Left Diagonal Win"
	InvalidMove   TurnResult = "Invalid Move"
	ValidMove     TurnResult = "Valid Move"
	Draw          TurnResult = "Draw"
)

func (r TurnResult) IsWin() bool {
	return r == HorizontalWin || r == VerticalWin || r == RDiagonalWin || r == LDiagonalWin
}
