package tree

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func CreateTreeView(cnames []string, minW []int) *gtk.TreeView {
	treev, _ := gtk.TreeViewNew()
	for i, cname := range cnames {
		renderer, _ := gtk.CellRendererTextNew()
		col, _ := gtk.TreeViewColumnNewWithAttribute(cname, renderer, "text", i)
		col.SetResizable(true)
		col.SetMinWidth(minW[i])
		treev.AppendColumn(col)
	}
	return treev
}

func CreateListStore(gtypes []glib.Type, data [][]interface{}) *gtk.ListStore {
	var nu []int
	for i := 0; i < len(gtypes); i++ {
		nu = append(nu, i)
	}
	store, _ := gtk.ListStoreNew(gtypes...)
	for _, row := range data {
		iter := store.Append()
		store.Set(iter, nu, row)
	}
	return store
}

func UpdateListStore(store *gtk.ListStore, gtypes []glib.Type, data [][]interface{}) *gtk.ListStore {
	var nu []int
	for i := 0; i < len(gtypes); i++ {
		nu = append(nu, i)
	}
	store.Clear()
	for i, row := range data {
		row[0] = i + 1
		iter := store.Append()
		store.Set(iter, nu, row)
	}
	return store
}
