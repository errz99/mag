package menu

import (
	"github.com/gotk3/gotk3/gtk"
)

var NoStr = "-"

// FillWithItems creates menu items from a slice of names and icons.
func FillWithItems(menu *gtk.Menu, items []*gtk.MenuItem, names []string, icons []string) {
	n := 0

	for i, name := range names {
		if name == NoStr {
			sep, _ := gtk.SeparatorMenuItemNew()
			sep.Show()
			menu.Append(sep)
			continue

		} else {
			if icons != nil && icons[i] != NoStr {
				lbl, _ := gtk.LabelNew(name)
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

			} else {
				items[n], _ = gtk.MenuItemNewWithLabel(name)
				items[n].Show()
				menu.Append(items[n])
			}

			n++
		}
	}
}

// FillWithItemsAndAccels creates menu items from a slice of names and icons
// including AccelGroup.
func FillWithItemsAndAccels(
	menu *gtk.Menu, items []*gtk.MenuItem,
	names []string, accels []string, agroup *gtk.AccelGroup) {

	n := 0
	for _, name := range names {
		if name == NoStr {
			sep, _ := gtk.SeparatorMenuItemNew()
			sep.Show()
			menu.Append(sep)
			continue
		} else {
			items[n].SetLabel(name)
			items[n].Show()

			if accels[n] != NoStr {
				key, modifier := gtk.AcceleratorParse(accels[n])
				if agroup != nil {
					items[n].AddAccelerator(
						"activate",
						agroup,
						key,
						modifier,
						gtk.ACCEL_VISIBLE,
					)
				}
			}

			menu.Append(items[n])
			n++
		}
	}
}
