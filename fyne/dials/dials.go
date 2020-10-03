package dials

import (
	//"sync"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func AboutDialog(text1, text2 string, win fyne.Window) {
	content1 := widget.NewLabelWithStyle(text1,
		fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	content2 := widget.NewLabelWithStyle(text2,
		fyne.TextAlignCenter, fyne.TextStyle{Monospace: true})

	box := widget.NewVBox(content1, content2)
	dialog.ShowCustom("About", "OK", box, win)
}

func TwoLabelsDialog(title, text1, text2 string, win fyne.Window) {
	content1 := widget.NewLabelWithStyle(text1,
		fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	content2 := widget.NewLabelWithStyle(text2,
		fyne.TextAlignLeading, fyne.TextStyle{Monospace: true})

	box := widget.NewVBox(content1, content2)
	dialog.ShowCustom(title, "OK", box, win)
}

func QuestionDialog(title, text string, win fyne.Window) int {
	var resp int
	//var wg sync.WaitGroup
	//wg.Add(1)

	confirmCall := func(response bool) {
		//defer wg.Done()

		if response {
			resp = -5
		} else {
			resp = -6
		}
	}

	cnf := dialog.NewConfirm(title, text, confirmCall, win)
	cnf.SetDismissText("Cancel")
	cnf.SetConfirmText("OK")
	cnf.Show()

	//wg.Wait()
	return resp
}

func EntryDialog(title, defTxt string, win fyne.Window) (int, string) {
	var resp int
	var text string
	//var wg sync.WaitGroup
	//wg.Add(1)

	content := widget.NewEntry()
	content.SetPlaceHolder(defTxt)

	content.OnChanged = func(t string) {
		text = t
	}

	call := func(message bool) {
		//defer wg.Done()

		if message {
			resp = -5
		} else {
			resp = -6
		}
	}

	vbox := widget.NewVBox(content, widget.NewLabel(""))
	dialog.ShowCustomConfirm(title, "OK", "Cancel", vbox, call, win)
	//wg.Wait()
	return resp, text
}

// Multi entries dialos (only 2 possible for now)
func EntriesDialog(title string, names, defTxt []string, win fyne.Window) (int, []string) {
	var resp int
	//var wg sync.WaitGroup
	//wg.Add(1)

	num := len(names)
	data := make([]string, num, num)
	form := &widget.Form{}

	// Entry 1
	entry1 := widget.NewPasswordEntry()
	form.Append(names[0], entry1)
	entry1.SetPlaceHolder(defTxt[0])
	entry1.OnChanged = func(text string) { data[0] = text }

	// Entry 2
	entry2 := widget.NewPasswordEntry()
	form.Append(names[1], entry2)
	entry2.SetPlaceHolder(defTxt[1])
	entry2.OnChanged = func(text string) { data[1] = text }

	call := func(message bool) {
		//defer wg.Done()

		if message {
			resp = -5
		} else {
			resp = -6
		}
	}

	vbox := widget.NewVBox(form, widget.NewLabel(""))
	dialog.ShowCustomConfirm(title, "OK", "Cancel", vbox, call, win)
	//wg.Wait()
	return resp, data
}

// Ask for a password, key, etc in a password entry
func AccessDialog(title, text string, win fyne.Window) (int, string) {
	var resp int
	var pass string
	//var wg sync.WaitGroup
	//wg.Add(1)

	content := widget.NewPasswordEntry()
	content.SetPlaceHolder(text)

	content.OnChanged = func(t string) {
		pass = t
	}

	call := func(message bool) {
		//defer wg.Done()

		if message {
			resp = -5
		} else {
			resp = -6
		}
	}

	vbox := widget.NewVBox(content, widget.NewLabel(""))
	dialog.ShowCustomConfirm(title, "OK", "Cancel", vbox, call, win)
	//wg.Wait()
	return resp, pass
}
