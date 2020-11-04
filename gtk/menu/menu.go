package menu

import (
	"github.com/gotk3/gotk3/gtk"
)

func FillWithItems(menu *gtk.Menu, items []*gtk.MenuItem, names []string, icons []string) {
	n := 0

	for i := 0; i < len(names); i++ {
		if names[i] == "-" {
			sep, _ := gtk.SeparatorMenuItemNew()
			sep.Show()
			menu.Append(sep)
			continue

		} else {
			if icons != nil && icons[i] != "-" {
				lbl, _ := gtk.LabelNew(names[i])
				icon, _ := gtk.ImageNew()
				icon.SetSizeRequest(16, 16)
				icon.SetFromIconName(icons[i], gtk.ICON_SIZE_MENU)
				box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 4)
				box.Add(icon)
				box.Add(lbl)
				items[n], _ = gtk.MenuItemNew()
				items[n].Add(box)

				items[n].Show()
				menu.Append(items[n])
				n++

			} else {
				items[n], _ = gtk.MenuItemNewWithLabel(names[i])
				items[n].Show()
				menu.Append(items[n])
				n++
			}
		}
	}
}
