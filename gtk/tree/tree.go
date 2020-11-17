package tree

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func CreateTreeView(cnames []string, minW []int) (*gtk.TreeView, error) {
	treev, _ := gtk.TreeViewNew()
	for i, cname := range cnames {
		if renderer, err := gtk.CellRendererTextNew(); err != nil {
			return nil, err
		} else {
			if col, err := gtk.TreeViewColumnNewWithAttribute(
				cname, renderer, "text", i); err != nil {
				return nil, err
			} else {
				col.SetResizable(true)
				col.SetMinWidth(minW[i])
				treev.AppendColumn(col)
			}
		}
	}
	return treev, nil
}

func CreateListStore(gtypes []glib.Type, data [][]interface{}) (*gtk.ListStore, error) {
	var nu []int
	for i := 0; i < len(gtypes); i++ {
		nu = append(nu, i)
	}
	store, err := gtk.ListStoreNew(gtypes...)
	if err != nil {
		return nil, err
	} else {
		for _, row := range data {
			iter := store.Append()
			store.Set(iter, nu, row)
		}
	}
	return store, nil
}

func UpdateListStore(store *gtk.ListStore, gtypes []glib.Type,
	data [][]interface{}) *gtk.ListStore {
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
