package main

import (
	"go-tic-tac-toe/grid"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var playgrid = grid.NewGrid(3)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")

	gameRunning := true

	label := widget.NewLabel("Tic Tac Toe")
	var resetBtn fyne.CanvasObject

	getTable := func(playgrid grid.Grid) *widget.Table {
		return widget.NewTable(
			func() (int, int) {
				return playgrid.Size(), playgrid.Size()
			},
			func() fyne.CanvasObject {
				btn := widget.NewButton("            ", func() {})
				return btn
			},
			func(i widget.TableCellID, o fyne.CanvasObject) {
				o.(*widget.Button).OnTapped = func() {
					if !gameRunning {
						return
					}

					res := playgrid.Mark(&grid.Position{X: i.Row, Y: i.Col})
					o.(*widget.Button).SetText(playgrid.GetMat()[i.Row][i.Col])

					if res.IsWin() {
						label.SetText("!! Victory !!")
						gameRunning = false

						resetBtn.Show()
					} else if res == grid.Draw {
						label.SetText("!! Draw !!")
						gameRunning = false

						resetBtn.Show()
					}
				}
			})
	}

	table := getTable(playgrid)
	var appLayout *fyne.Container

	resetBtn = widget.NewButton("Reset", func() {
		playgrid.Reset()
		table = getTable(playgrid)
		appLayout.Objects[1] = table

		gameRunning = true
		label.SetText("Tic Tac Toe")
		resetBtn.Hide()
	})
	resetBtn.Hide()

	header := container.New(layout.NewGridLayout(2), label, resetBtn)
	appLayout = container.New(layout.NewGridLayout(1), header, table)

	myWindow.SetContent(appLayout)
	myWindow.Resize(fyne.NewSize(200, 250))
	myWindow.ShowAndRun()
}
