package main

type CellType string

const (
	Empty    CellType = "-"
	CellX    CellType = "X"
	CellO    CellType = "O"
	CellDraw CellType = "V"
)

func (c CellType) String() string {
	switch c {
	case CellO:
		return ColorGreen.Paint(string(c))
	case CellX:
		return ColorPurple.Paint(string(c))
	case CellDraw:
		return ColorRed.Paint(string(c))
	default:
		return string(c)
	}
}
