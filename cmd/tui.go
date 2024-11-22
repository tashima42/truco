package cmd

import "github.com/rivo/tview"

func Run() error {
	// box := tview.NewBox().SetBorder(true).SetTitle("Truco")
	// app := tview.NewApplication().SetRoot(box, true)
	//
	// button := tview.NewButton("Quit").SetSelectedFunc(func() {
	// 	app.Stop()
	// })
	// button.SetBorder(true).SetRect(0, 0, 22, 3)
	// if err := app.SetRoot(button, false).SetFocus(button).Run(); err != nil {
	// 	panic(err)
	// }
	//
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	header := newPrimitive("Truco")
	menu := newPrimitive("Menu")
	main := newPrimitive("Main content")

	grid := tview.NewGrid().
		SetRows(2, 0).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(header, 0, 0, 1, 3, 0, 0, false).
		AddItem(menu, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 2, 0, 100, false)

	app := tview.NewApplication()

	button := tview.NewButton("Quit").SetSelectedFunc(func() {
		app.Stop()
	})
	button.SetBorder(false).SetRect(0, 0, 10, 5)
	// show button on the menu
	grid.AddItem(button, 1, 0, 1, 1, 0, 0, false)
	// add a button on the menu grid to quit the application

	return app.SetRoot(grid, true).SetFocus(grid).Run()
}
