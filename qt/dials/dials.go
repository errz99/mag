package dials

// Last edit: 20200103

import (
	//"fmt"
	//"os"

	"github.com/errz99/mag"
	"github.com/therecipe/qt/core"
	//"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

//var md1a = "<span lang=\"utf-8\" color=\"green\"><b>"
//var md1b =	"</b></span>"
//
//var md2a = "<span lang=\"utf-8\"><tt>"
//var md2b =	"</tt></span>"

// Accept ...
var Accept = -5

// Cancel ...
var Cancel = -6

// Initial language set
var lang mag.Language = mag.Lang

func SetLanguage(lng string) {
	switch strings.ToLower(lng) {
	case "es":
		lang = mag.Spanish
	case "en":
		lang = mag.English
	default:
	}
}

// Menu Help
func AboutDialog(title, text string, parent *widgets.QMainWindow) {
	widgets.QMessageBox_About(parent, title, text)
}

// Message IWE, for Info, Warning, Error, "", dialogs
func MessageIWE(w, h int, class, title, msg, info string, win *widgets.QMainWindow) {
	mbox := widgets.NewQMessageBox(win)

	switch class {
	case "Info":
		mbox.SetIcon(widgets.QMessageBox__Information)

	case "Question":
		mbox.SetIcon(widgets.QMessageBox__Question)
		mbox.AddButton3(widgets.QMessageBox__Cancel)
		mbox.AddButton3(widgets.QMessageBox__Ok)
		mbox.SetDefaultButton2(widgets.QMessageBox__Cancel)

	case "Warning":
		mbox.SetIcon(widgets.QMessageBox__Warning)

	case "Error":
		mbox.SetIcon(widgets.QMessageBox__Critical)

	default:
		mbox.SetIcon(widgets.QMessageBox__NoIcon)
	}

	mbox.SetWindowTitle(title)
	mbox.SetText(msg)
	mbox.SetInformativeText(info)
	mbox.Show()
	mbox.Exec()
}

// MessageIWED ...
func MessageIWED(w, h int, class, title, msg, info string, win *widgets.QDialog) {
	mbox := widgets.NewQMessageBox(win)

	switch class {
	case "Info":
		mbox.SetIcon(widgets.QMessageBox__Information)

	case "Question":
		mbox.SetIcon(widgets.QMessageBox__Question)
		mbox.AddButton3(widgets.QMessageBox__Cancel)
		mbox.AddButton3(widgets.QMessageBox__Ok)
		mbox.SetDefaultButton2(widgets.QMessageBox__Cancel)

	case "Warning":
		mbox.SetIcon(widgets.QMessageBox__Warning)

	case "Error":
		mbox.SetIcon(widgets.QMessageBox__Critical)

	default:
		mbox.SetIcon(widgets.QMessageBox__NoIcon)
	}

	mbox.SetWindowTitle(title)
	mbox.SetText(msg)
	mbox.SetInformativeText(info)
	mbox.Show()
	mbox.Exec()
}

// MessageQuestion dialog
func MessageQuestion(w, h int, title, mess string, win *widgets.QMainWindow) int {
	dialog := widgets.NewQMessageBox(win)
	answer := Cancel

	dialog.SetWindowTitle("Question")
	dialog.SetText(title)
	dialog.SetInformativeText(mess)

	dialog.SetIcon(widgets.QMessageBox__Question)
	dialog.AddButton3(widgets.QMessageBox__Cancel)
	dialog.AddButton3(widgets.QMessageBox__Ok)
	dialog.SetDefaultButton2(widgets.QMessageBox__Cancel)

	dialog.ConnectButtonClicked(func(button *widgets.QAbstractButton) {
		if button.Text() == lang["ButtonOK"] {
			answer = Accept
		}
	})

	dialog.Show()
	answer = dialog.Exec()

	if answer == 1024 {
		answer = Accept
	} else {
		answer = Cancel
	}

	return answer
}

// DialogTwoLabels for two labels
func DialogTwoLabels(title, text1, text2 string, win *widgets.QMainWindow) {
	dialog := widgets.NewQDialog(win, 0)
	dialog.SetModal(true)
	dialog.SetWindowTitle(title)
	dialog.SetMinimumSize2(200, 120)

	vbox := widgets.NewQVBoxLayout()
	dialog.SetLayout(vbox)

	label1 := widgets.NewQLabel2(text1, dialog, 0)
	label1.SetStyleSheet("color: green;")

	label2 := widgets.NewQLabel2(text2, dialog, 0)

	vbox.AddWidget(label1, 0, core.Qt__AlignCenter)
	vbox.AddWidget(label2, 0, core.Qt__AlignCenter)

	okButton := widgets.NewQPushButton2("Ok", dialog)
	vbox.AddWidget(okButton, 0, core.Qt__AlignRight)

	okButton.ConnectClicked(func(bool) {
		dialog.DestroyQDialog()
	})

	dialog.Show()
	dialog.Exec()
}

// DialogEntry for an entry alone
func DialogEntry(title, head, entryLabel, entryPre string,
	visible bool, win *widgets.QMainWindow) (int, string) {
	dialog := widgets.NewQDialog(win, 0)
	dialog.SetModal(true)
	dialog.SetWindowTitle(title)
	dialog.SetMinimumSize2(200, 120)

	vbox := widgets.NewQVBoxLayout()
	dialog.SetLayout(vbox)

	entry := widgets.NewQWidget(dialog, 0)
	form := widgets.NewQFormLayout(entry)
	entry.SetLayout(form)

	label := widgets.NewQLabel2(head, dialog, 0)
	label.SetStyleSheet("color: green;")

	buttons := widgets.NewQWidget(dialog, 0)
	hbox := widgets.NewQHBoxLayout()
	buttons.SetLayout(hbox)

	vbox.AddWidget(label, 0, core.Qt__AlignCenter)
	vbox.AddWidget(entry, 0, core.Qt__AlignCenter)
	vbox.AddWidget(buttons, 0, core.Qt__AlignRight)

	lineEdit := widgets.NewQLineEdit(nil)
	lineEdit.SetPlaceholderText(entryPre)
	if !visible {
		lineEdit.SetEchoMode(widgets.QLineEdit__Password)
	}
	form.AddRow3(entryLabel, lineEdit)

	cancelButton := widgets.NewQPushButton2(lang["ButtonCancel"], dialog)
	hbox.AddWidget(cancelButton, 0, core.Qt__AlignRight)
	okButton := widgets.NewQPushButton2(lang["ButtonOK"], dialog)
	hbox.AddWidget(okButton, 0, core.Qt__AlignRight)

	var answer = Cancel
	var entryText = entryPre

	okButton.ConnectClicked(func(bool) {
		answer = Accept
		if lineEdit.Text() != "" {
			entryText = lineEdit.Text()
		}
		dialog.DestroyQDialog()
	})

	cancelButton.ConnectClicked(func(bool) {
		dialog.DestroyQDialog()
	})

	dialog.Show()
	dialog.Exec()

	return answer, entryText
}

// DialogEntries for multiple entries
func DialogEntries(title, head string, entryLabels, entryTexts []string,
	visibles []bool, win *widgets.QMainWindow) (int, []string) {
	dialog := widgets.NewQDialog(win, 0)
	dialog.SetModal(true)
	dialog.SetWindowTitle(title)

	vbox := widgets.NewQVBoxLayout()
	dialog.SetLayout(vbox)

	entry := widgets.NewQWidget(dialog, 0)
	form := widgets.NewQFormLayout(entry)
	entry.SetLayout(form)

	label := widgets.NewQLabel2(head, dialog, 0)
	label.SetStyleSheet("color : green;")

	buttons := widgets.NewQWidget(dialog, 0)
	hbox := widgets.NewQHBoxLayout()
	buttons.SetLayout(hbox)

	vbox.AddWidget(label, 0, core.Qt__AlignCenter)
	vbox.AddWidget(entry, 0, core.Qt__AlignCenter)
	vbox.AddWidget(buttons, 0, core.Qt__AlignRight)
	var lineEdits []*widgets.QLineEdit

	for i, entryText := range entryTexts {
		lineEdit := widgets.NewQLineEdit(nil)
		if !visibles[i] {
			lineEdit.SetEchoMode(widgets.QLineEdit__Password)
		}
		lineEdit.SetText(entryText)
		form.AddRow3(entryLabels[i], lineEdit)
		lineEdits = append(lineEdits, lineEdit)
	}

	cancelButton := widgets.NewQPushButton2(lang["ButtonCancel"], dialog)
	hbox.AddWidget(cancelButton, 0, core.Qt__AlignRight)
	okButton := widgets.NewQPushButton2(lang["ButtonOK"], dialog)
	hbox.AddWidget(okButton, 0, core.Qt__AlignRight)

	var answer = Cancel
	var lineTexts []string

	okButton.ConnectClicked(func(bool) {
		answer = Accept
		for _, lineEdit := range lineEdits {
			lineTexts = append(lineTexts, lineEdit.Text())
		}
		dialog.DestroyQDialog()
	})

	cancelButton.ConnectClicked(func(bool) {
		dialog.DestroyQDialog()
	})

	dialog.Show()
	dialog.Exec()

	return answer, lineTexts
}

// AccessKeyDialog ask for a password or access key.
func AccessKeyDialog(title, head string, win *widgets.QMainWindow) (int, string) {
	dialog := widgets.NewQDialog(win, 0)
	dialog.SetModal(true)
	dialog.SetWindowTitle(title)
	dialog.SetMinimumSize2(200, 1)

	vbox := widgets.NewQVBoxLayout()
	dialog.SetLayout(vbox)

	entry := widgets.NewQWidget(dialog, 0)
	form := widgets.NewQFormLayout(entry)
	entry.SetLayout(form)

	if head != "" {
		label := widgets.NewQLabel2(head, dialog, 0)
		label.SetStyleSheet("color: green;")
		vbox.AddWidget(label, 0, core.Qt__AlignCenter)
	}

	buttons := widgets.NewQWidget(dialog, 0)
	hbox := widgets.NewQHBoxLayout()
	buttons.SetLayout(hbox)

	//vbox.AddWidget(label, 0, core.Qt__AlignCenter)
	vbox.AddWidget(entry, 0, core.Qt__AlignCenter)
	vbox.AddWidget(buttons, 0, core.Qt__AlignRight)

	lineEdit := widgets.NewQLineEdit(nil)
	lineEdit.SetMinimumSize2(200, 1)
	lineEdit.SetAlignment(core.Qt__AlignCenter)
	lineEdit.SetEchoMode(widgets.QLineEdit__Password)
	form.AddRow3("", lineEdit)

	okButton := widgets.NewQPushButton2(lang["ButtonOK"], dialog)
	hbox.AddWidget(okButton, 0, core.Qt__AlignRight)

	var answer = Cancel
	var entryText string

	okButton.ConnectClicked(func(bool) {
		answer = Accept
		entryText = lineEdit.Text()
		dialog.DestroyQDialog()
	})

	dialog.Show()
	dialog.Exec()

	return answer, entryText
}

// FileDialogForOpen returns a file path selected from of a dialog
func FileDialogForOpen(parent widgets.QWidget_ITF, caption, dir, filter string) (int, string) {
	fileDialog := widgets.NewQFileDialog2(parent, caption, dir, filter)
	fileDialog.SetFileMode(widgets.QFileDialog__AnyFile)

	resp := fileDialog.Exec()
	if resp == 1 && fileDialog.SelectedNameFilter() == filter {
		if len(fileDialog.SelectedFiles()) == 1 {
			return resp, fileDialog.SelectedFiles()[0]
		}
	}

	return 0, ""
}

// FileDialogForSave returns a file path selected from of a dialog
func FileDialogForSave(parent widgets.QWidget_ITF, caption, dir, filter string) (int, string) {
	fileDialog := widgets.NewQFileDialog2(parent, caption, dir, filter)
	fileDialog.SetAcceptMode(widgets.QFileDialog__AcceptSave)

	resp := fileDialog.Exec()
	if resp == 1 && fileDialog.SelectedNameFilter() == filter {
		if len(fileDialog.SelectedFiles()) == 1 {
			return resp, fileDialog.SelectedFiles()[0]
		}
	}

	return 0, ""
}
