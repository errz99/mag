package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/errz99/mag/gtk/tree"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(nil)
	mainWindow()
	gtk.Main()

}

func mainWindow() {
	cL := "|    "
	cR := "    |"

	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Tree View")

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
		"Great Musicians",
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
				resp, sel, name := musiciansTree(window)
				if resp != -6 {
					fmt.Println("Musician selected:", sel, name)
				}

			case "2":
				gtk.MainQuit()

			default:
			}
		})
	}

	window.ShowAll()
}

func musiciansTree(win gtk.IWindow) (gtk.ResponseType, int, string) {
	dial, _ := gtk.DialogNew()
	dial.SetTransientFor(win)
	dial.SetTitle("Great Musicians")

	content, _ := dial.GetContentArea()
	content.SetBorderWidth(8)
	content.SetSpacing(4)
	label, _ := gtk.LabelNew("Rock Musicians")

	content.PackStart(label, false, true, 4)

	// Tree View
	titles := []string{"Nr", "First Name", "Last Name"}
	minWidth := []int{48, 96, 96}
	gtypes := []glib.Type{glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING}

	data := [][]interface{}{
		{1, "Bob", "Dylan"},
		{2, "Ian", "Anderson"},
		{3, "Jorma", "Kaukonen"},
		{4, "Jerry", "Garcia"},
	}

	treeview, _ := tree.CreateTreeView(titles, minWidth)
	store, _ := tree.CreateListStore(gtypes, data)

	treeview.SetModel(store)
	treeview.SetVExpand(true)
	treeview.ColumnsAutosize()
	treeview.SetMarginBottom(8)
	content.Add(treeview)

	var presel string
	var path *gtk.TreePath
	var selected int
	var name string

	if presel != "" {
		iter, _ := store.GetIterFromString(presel)
		tsel, _ := treeview.GetSelection()
		tsel.SelectIter(iter)
	}

	selectMusicianName := func() {
		path, _ = treeview.GetCursor()
		ps := path.String()
		if ps != "" {
			selected, _ = strconv.Atoi(ps)
			iter, _ := store.GetIter(path)
			first, _ := store.GetValue(iter, 1)
			last, _ := store.GetValue(iter, 2)
			firstName, _ := first.GetString()
			lastName, _ := last.GetString()
			name = firstName + " " + lastName
		}
	}

	treeview.Connect("row-activated", func() {
		selectMusicianName()
		dial.Destroy()
	})

	// Buttons
	dial.AddButton("Cancel", -6)
	dial.AddButton("OK", -5)
	dial.SetDefaultResponse(-5)
	dial.ShowAll()

	answer := dial.Run()
	if answer != -1 {
		selectMusicianName()
	}
	dial.Destroy()

	return answer, selected, name
}
