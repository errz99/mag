package dials

import (
	"github.com/gotk3/gotk3/gtk"
)

// Message shows an Info, Warning or Error message
func Message(w, h int, class, title, msg string, win gtk.IWindow) {
	var dialog *gtk.MessageDialog
	switch class {
	case "Info":
		dialog = gtk.MessageDialogNew(win, 1, gtk.MESSAGE_INFO, gtk.BUTTONS_CLOSE, msg)
	case "Warning":
		dialog = gtk.MessageDialogNew(win, 1, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, msg)
	case "Error":
		dialog = gtk.MessageDialogNew(win, 1, gtk.MESSAGE_ERROR, gtk.BUTTONS_CLOSE, msg)
	default:
		dialog = gtk.MessageDialogNew(win, 1, gtk.MESSAGE_OTHER, gtk.BUTTONS_CLOSE, msg)
	}

	dialog.SetTitle(title)
	dialog.SetSizeRequest(w, h)
	dialog.Run()
	dialog.Destroy()
}

// Question shows a dialog allowing to cancel or accept a question
func Question(w, h int, title, msg string, win gtk.IWindow) gtk.ResponseType {
	dialog := gtk.MessageDialogNew(win, 1, gtk.MESSAGE_QUESTION, gtk.BUTTONS_OK_CANCEL, msg)
	dialog.SetTitle(title)
	dialog.SetSizeRequest(w, h)
	answer := dialog.Run()
	dialog.Destroy()

	return answer
}

// OneEntry shows a dialog with an entry widget
func OneEntry(title, head, entryLabel, entryText string,
	visible bool, win gtk.IWindow) (gtk.ResponseType, string) {

	dial, _ := gtk.DialogNew()
	dial.SetTransientFor(win)
	dial.SetTitle(title)

	content, _ := dial.GetContentArea()
	content.SetSpacing(4)
	label, _ := gtk.LabelNew(head)

	hbox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 4)
	name, _ := gtk.LabelNew(entryLabel)
	entry, _ := gtk.EntryNew()
	entry.SetText(entryText)
	entry.SetWidthChars(32)
	entry.SetVisibility(visible)
	hbox.PackStart(name, false, true, 4)
	hbox.PackStart(entry, true, true, 4)

	content.PackStart(label, false, true, 4)
	content.PackStart(hbox, false, true, 4)

	dial.AddButton("Cancel", -6)
	dial.AddButton("OK", -5)
	dial.SetDefaultResponse(-5)
	dial.ShowAll()

	answer := dial.Run()
	entrName, _ := entry.GetText()

	dial.Destroy()
	return answer, entrName
}

// MultiEntries shows a dialog with multiple entries
func MultiEntries(title, head string, entryLabels, entryTexts []string,
	visibles []bool, win gtk.IWindow) (gtk.ResponseType, []string) {

	dial, _ := gtk.DialogNew()
	dial.SetTransientFor(win)
	dial.SetTitle(title)

	content, _ := dial.GetContentArea()
	content.SetSpacing(6)

	label, _ := gtk.LabelNew(head)
	content.PackStart(label, false, true, 4)
	var entries []*gtk.Entry

	grid, _ := gtk.GridNew()
	grid.SetHAlign(gtk.ALIGN_FILL)
	grid.SetRowSpacing(6)
	grid.SetColumnSpacing(6)
	grid.SetMarginStart(6)
	grid.SetMarginEnd(6)
	content.Add(grid)

	for i, entryText := range entryTexts {
		name, _ := gtk.LabelNew(entryLabels[i])
		name.SetHAlign(gtk.ALIGN_END)
		entry, _ := gtk.EntryNew()
		entry.SetText(entryText)
		entry.SetWidthChars(32)
		entry.SetVisibility(visibles[i])
		entries = append(entries, entry)

		grid.Attach(name, 0, i, 1, 1)
		grid.Attach(entry, 1, i, 1, 1)
	}

	dial.AddButton("Cancel", -6)
	dial.AddButton("OK", -5)
	dial.SetDefaultResponse(-5)
	dial.ShowAll()

	answer := dial.Run()
	var entryNames []string

	for _, entry := range entries {
		text, _ := entry.GetText()
		entryNames = append(entryNames, text)
	}

	dial.Destroy()
	return answer, entryNames
}

// OnePassword returns a password (response -5 for Ok button and -1 for activate event)
func OnePassword(title string, win gtk.IWindow) (gtk.ResponseType, string) {
	dial, _ := gtk.DialogNew()
	dial.SetTransientFor(win)
	dial.SetTitle(title)

	content, _ := dial.GetContentArea()
	content.SetSpacing(4)

	hbox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 4)
	entry, _ := gtk.EntryNew()
	entry.SetWidthChars(32)
	entry.SetVisibility(false)
	entry.SetAlignment(0.5)
	hbox.PackStart(entry, true, true, 4)

	content.PackStart(hbox, false, true, 4)

	dial.AddButton("OK", -5)
	dial.SetDefaultResponse(-6)
	dial.ShowAll()

	var answer gtk.ResponseType = -6
	var entrName string

	entry.Connect("changed", func() {
		entrName, _ = entry.GetText()
	})

	entry.Connect("activate", func() {
		dial.Destroy()
	})

	answer = dial.Run()
	dial.Destroy()
	return answer, entrName
}

// ChooseAFile returns a file name (with path included)
func ChooseAFileForOpen(title, current string, win gtk.IWindow) (gtk.ResponseType, string, error) {
	fchooser, err := gtk.FileChooserDialogNewWith2Buttons(
		title, win, gtk.FILE_CHOOSER_ACTION_OPEN, "Cancel", -6, "Open", -5)

	if err != nil {
		return 0, "", err
	}

	fchooser.SetCurrentFolder(current)
	response := fchooser.Run()
	filename := fchooser.GetFilename()
	fchooser.Destroy()
	return response, filename, nil
}

// ChooseAFileForSave returns a file name (with path included)
func ChooseAFileForSave(title, current string, win gtk.IWindow) (gtk.ResponseType, string, error) {
	fchooser, err := gtk.FileChooserDialogNewWith2Buttons(
		title, win, gtk.FILE_CHOOSER_ACTION_SAVE, "Cancel", -6, "Save", -5)

	if err != nil {
		return 0, "", err
	}

	fchooser.SetCurrentFolder(current)
	response := fchooser.Run()
	filename := fchooser.GetFilename()
	fchooser.Destroy()
	return response, filename, nil
}

// ShowEditText shows a text in a TextView and returns its edition
func ShowEditText(title, head, file, text string, parent gtk.IWindow) (gtk.ResponseType, string) {
	dial, _ := gtk.DialogNew()
	dial.SetTransientFor(parent)
	dial.SetTitle(title)
	dial.SetDefaultSize(0, 0)

	content, _ := dial.GetContentArea()
	content.SetSpacing(4)
	content.SetBorderWidth(8)

	label, _ := gtk.LabelNew(head)
	label.SetMarkup("<span foreground=\"red\" size=\"medium\">" + head + "</span>")
	label.SetMarginBottom(4)
	content.Add(label)

	tag, _ := gtk.TextTagNew("tag1")
	table, _ := gtk.TextTagTableNew()
	table.Add(tag)
	buffer, _ := gtk.TextBufferNew(table)
	buffer.SetText(text)
	tview, _ := gtk.TextViewNewWithBuffer(buffer)

	tview.SetLeftMargin(4)
	tview.SetRightMargin(4)
	tview.SetWrapMode(gtk.WRAP_WORD)
	content.Add(tview)

	dial.AddButton("Cancel", -6)
	dial.AddButton("OK", -5)
	dial.SetDefaultResponse(-6)
	dial.ShowAll()

	answer := dial.Run()
	a, b := buffer.GetBounds()
	text, _ = buffer.GetText(a, b, false)

	dial.Destroy()
	return answer, text
}

// TwoLabels shows two texts with different colors
func TwoLabels(title, text1, text2 string, parent gtk.IWindow) {
	dial, _ := gtk.DialogNew()
	dial.SetTransientFor(parent)
	dial.SetTitle(title)

	content, _ := dial.GetContentArea()
	content.SetSpacing(6)
	content.SetMarginTop(6)
	content.SetMarginBottom(6)
	content.SetMarginStart(12)
	content.SetMarginEnd(12)

	label1, _ := gtk.LabelNew(text1)
	label1.SetMarkup("<span color=\"green\"><b>" + text1 + "</b></span>")
	label2, _ := gtk.LabelNew(text2)
	label2.SetMarkup("<span><tt>" + text2 + "</tt></span>")
	content.Add(label1)
	content.Add(label2)

	dial.AddButton("OK", -5)
	dial.ShowAll()

	dial.Run()
	dial.Destroy()
}
