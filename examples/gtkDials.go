package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/errz99/mag/gtk/dials"
	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(nil)
	mainWindow()
	gtk.Main()

}

var home = os.Getenv("HOME")

func mainWindow() {
	cL := "|    "
	cR := "    |"

	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Dialogs")

	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	vbox.SetBorderWidth(5)
	vbox.SetCanFocus(false)
	window.Add(vbox)

	head, _ := gtk.LabelNew("")
	head.SetMarkup("<span color=\"green\"><b>Select a Dialog</b></span>")
	head.SetMarginTop(8)
	head.SetMarginBottom(8)
	vbox.Add(head)

	buttonTexts := []string{
		"Info, Warning, Error",
		"Question Dialog",
		"Two Labels",
		"Open a File",
		"Save a File",
		"One Entry",
		"Multi Entries",
		"Show Edit Text",
		"Exit"}

	for i, text := range buttonTexts {
		button, _ := gtk.ButtonNewWithLabel(text)
		button.SetName(strconv.Itoa(i + 1))
		button.SetRelief(gtk.RELIEF_HALF)
		button.SetSizeRequest(240, 0)
		vbox.Add(button)

		button.Connect("focus-in-event", func() {
			label, _ := button.GetLabel()
			button.SetLabel(cL + label + cR)
		})

		button.Connect("focus-out-event", func() {
			label, _ := button.GetLabel()
			label = strings.TrimLeft(label, cL)
			label = strings.TrimRight(label, cR)
			button.SetLabel(label)
		})

		button.Connect("clicked", func() {
			name, _ := button.GetName()

			switch name {
			case "1":
				dials.Message(0, 0, "Info", "Message Dialog",
					"Info | Warning | Error Message", window)

			case "2":
				dials.Question(0, 0, "Question Dialog", "Wanna continue?", window)

			case "3":
				dials.TwoLabels("Title", "Text One", "Text Two", window)

			case "4":
				resp, file, _ := dials.ChooseAFileForOpen("Open File", home, window)
				if resp == -5 {
					dials.Message(0, 0, "Info", "File Selected for Open", file, window)
				}

			case "5":
				resp, file, _ := dials.ChooseAFileForSave("Save File", home, window)
				if resp == -5 {
					dials.Message(0, 0, "Info", "File Selected for Save", file, window)
				}

			case "6":
				resp, text := dials.OneEntry("Title", "Head", "Label", "Entry text", true, window)
				if resp == -5 {
					dials.Message(0, 0, "Info", "Entry text", text, window)
				}

			case "7":
				resp, texts := dials.MultiEntries("Title", "Head",
					[]string{"Label1", "Label2"},
					[]string{"Entry1 text", "Entry2 text"},
					[]bool{true, true}, window)
				if resp == -5 {
					dials.Message(0, 0, "Info", "Entries texts", texts[0]+"\n...", window)
				}

			case "8":
				editable := "This text can be edited\nand returned."
				resp, edited := dials.ShowEditText("Title", "Head",
					"file text here\n and here", editable, window)
				if resp == -5 {
					dials.Message(0, 0, "Info", "Edited Text", edited, window)
				}

			case "9":
				gtk.MainQuit()

			default:
			}
		})
	}

	window.ShowAll()
}
