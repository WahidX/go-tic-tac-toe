package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{
	{"x", "x", "x"},
	{"x", "x", "x"},
	{"x", "x", "x"},
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewButton("wide content", func() {
				fmt.Println("tapped")
			})
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Button).SetText(data[i.Row][i.Col])
			o.(*widget.Button).OnTapped = func() {
				fmt.Println("tapped", i)

				o.(*widget.Button).SetText("tapped")
			}
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}
