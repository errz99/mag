package dials

// Last edit: 20200103

import (
	//"fmt"
    //"os"

	"github.com/therecipe/qt/core"
	//"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"	
)


// Menu Help
func aboutDialog() {
	//widgets.QMessageBox.About()

	// about, _ := gtk.AboutDialogNew()
 //    about.SetTransientFor(mwin)
 //    about.SetAuthors([]string{"M Arias"})
 //    about.SetComments("Un comentario aqui")
 //    about.SetCopyright("Copyryght aqui")
 //    about.SetProgramName("Pasun")
 //    about.SetVersion(VERSION)
 //    about.SetWebsite("www.arteop.com")

	// about.ShowAll()
	// about.Run()
	// about.Destroy()
}

//var md1a = "<span lang=\"utf-8\" color=\"green\"><b>"
//var md1b =	"</b></span>"
//
//var md2a = "<span lang=\"utf-8\"><tt>"
//var md2b =	"</tt></span>"

// Info, Warning, Error, ""
func MessageIWE(w, h int, class, title, mess string, win *widgets.QMainWindow) {
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

    mbox.SetWindowTitle(class)
    mbox.SetText(title)
    mbox.SetInformativeText(mess)
    mbox.Show()
}

func MessageIWED(w, h int, class, title, mess string, win *widgets.QDialog) {
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

    mbox.SetWindowTitle(class)
    mbox.SetText(title)
    mbox.SetInformativeText(mess)
    mbox.Show()
}

// Question dialog
func MessageQuestion(w, h int, title, mess string, win *widgets.QMainWindow) int {
    dialog := widgets.NewQMessageBox(win)
    answer := -6

    dialog.SetWindowTitle("Question")
    dialog.SetText(title)
    dialog.SetInformativeText(mess)

    dialog.SetIcon(widgets.QMessageBox__Question)
    dialog.AddButton3(widgets.QMessageBox__Cancel)
    dialog.AddButton3(widgets.QMessageBox__Ok)
    dialog.SetDefaultButton2(widgets.QMessageBox__Cancel)

	dialog.ConnectButtonClicked(func(button *widgets.QAbstractButton) {
		if button.Text() == "Ok" {
			answer = -5
		}
	})

    dialog.Show()
    answer = dialog.Exec()
    
    if answer == 1024 {
    	answer = -5 
    } else {
    	 answer = -6 
    }

    return answer
}

// Dialog for two labels
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
}

// Dialog for an entry alone
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
	form.AddRow3(entryLabel, lineEdit)

	cancelButton := widgets.NewQPushButton2("Cancel", dialog)
	hbox.AddWidget(cancelButton, 0, core.Qt__AlignRight)
	okButton := widgets.NewQPushButton2("Ok", dialog)
	hbox.AddWidget(okButton, 0, core.Qt__AlignRight)

	var answer = -6
	var entryText = entryPre

	okButton.ConnectClicked(func(bool) {
		answer = -5
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

// Dialog for multiple entries
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
		if visibles[i] == true {
			lineEdit.SetEchoMode(widgets.QLineEdit__Password)
		}
		lineEdit.SetText(entryText)
		form.AddRow3(entryLabels[i], lineEdit)
		lineEdits = append(lineEdits, lineEdit)
	}	

	cancelButton := widgets.NewQPushButton2("Cancel", dialog)
	hbox.AddWidget(cancelButton, 0, core.Qt__AlignRight)
	okButton := widgets.NewQPushButton2("Ok", dialog)
	hbox.AddWidget(okButton, 0, core.Qt__AlignRight)

	var answer = -6
	var lineTexts []string

	okButton.ConnectClicked(func(bool) {
		answer = -5
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

// Dialog that returns a file name (and path)
//func GetAFileDialog(head string, win *gtk.Window) (gtk.ResponseType, string) {
//	fchooser, _ := gtk.FileChooserDialogNewWith2Buttons(head, win,
//		gtk.FILE_CHOOSER_ACTION_OPEN, "Open", -5, "Cancel", -6)
//	fchooser.SetCurrentFolder(os.Getenv("HOME") + "/")
//	resp := fchooser.Run()
//	fname := fchooser.GetFilename()
//	fchooser.Destroy()
//	return resp, fname
//}
