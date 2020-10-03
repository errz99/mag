package main

import (
	"fmt"
	//"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/errz99/mag/fyne/dials"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())

	mwin := a.NewWindow("Fyne Checks")
	mainWin(a, mwin)
}

func mainWin(a fyne.App, mwin fyne.Window) {

	content := fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		widget.NewLabelWithStyle("Select Dialog",
			fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),

		widget.NewButton("About Dialog", func() {
			dials.AboutDialog("Text One", "Text Two", mwin)
		}),

		widget.NewButton("Question Dialog", func() {
			resp := dials.QuestionDialog("Title", "Wanna leave?", mwin)
			fmt.Println(resp)

		}),

		widget.NewButton("Entry Dialog", func() {
			resp, text := dials.EntryDialog("Title", "Default Text", mwin)
			fmt.Println(resp, text)
		}),

		widget.NewButton("Two Entries Dialog", func() {
			resp, texts := dials.EntriesDialog("Title",
				[]string{"Name One", "Name Two"},
				[]string{"Def Text One", "Def Text Two"}, mwin)
			fmt.Println(resp, texts)
		}),

		widget.NewButton("Access Dialog", func() {
			resp, text := dials.AccessDialog("Title", "This is The Text", mwin)
			fmt.Println(resp, text)
		}),

		widget.NewButton("Show Confirm", func() {
			dialog.ShowConfirm("Title", "Do you wanna dance?",
			func(resp bool) {
				if resp {
					fmt.Println("Yes, I do")
				} else {
					fmt.Println("No, I don't")
				}		
			},
			mwin)			
		}),

		widget.NewButton("Info Dialog", func() {
			infoDial := dialog.NewInformation("Title", "Very informative text", mwin)
			infoDial.Show()
		}),

		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	)

	mwin.SetContent(content)

	mwin.SetOnClosed(func() {
		//saveBeforeLeave(model)
		//a.Quit()
		//fmt.Println("CERRANDOOOO...")
	})
	mwin.Resize(fyne.Size{400, 1})

	mwin.ShowAndRun()
}
