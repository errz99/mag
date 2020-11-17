package mag

import (
//"fmt"
)

type Language map[string]string

var English Language = Language{
	// for Gtk
	"ButtonOK":     "_OK",
	"ButtonCancel": "_Cancel",
	"ButtonOpen":   "_Open",
	"ButtonSave":   "_Save",
	// for Qt
	"QButtonOK":     "&OK",
	"QButtonCancel": "&Cancel",
	"QButtonOpen":   "&Open",
	"QButtonSave":   "&Save",
}

var Spanish Language = Language{
	// for Gtk
	"ButtonOK":     "_Aceptar",
	"ButtonCancel": "_Cancelar",
	"ButtonOpen":   "_Abrir",
	"ButtonSave":   "_Guardar",
	// for Qt
	"QButtonOK":     "&Aceptar",
	"QButtonCancel": "&Cancelar",
	"QButtonOpen":   "&Abrir",
	"QButtonSave":   "&Guardar",
}

var Lang Language = English

func Hello() string {
	return "Hello, world."
}
