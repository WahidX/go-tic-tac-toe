package main

import (
	"go-snake/grid"

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

	table := widget.NewTable(
		func() (int, int) {
			return playgrid.Size(), playgrid.Size()
		},
		func() fyne.CanvasObject {
			btn := widget.NewButton(" ", func() {})
			btn.Resize(fyne.NewSize(200, 200))
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
					return
				}
			}
		})

	appLayout := container.New(layout.NewGridLayoutWithRows(2), label, table)

	myWindow.SetContent(appLayout)
	myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.ShowAndRun()
}
