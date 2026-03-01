package gui

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var (
	sidebarData = map[string]func(*Config, fyne.App, fyne.Window) fyne.CanvasObject{
		"Company": func(c *Config, a fyne.App, w fyne.Window) fyne.CanvasObject {
			return c.CompanyPage(a, w)
		},
		"Payment Method": func(c *Config, a fyne.App, w fyne.Window) fyne.CanvasObject {
			return c.PaymentDetailPage(a, w)
		},
		"Item": func(c *Config, a fyne.App, w fyne.Window) fyne.CanvasObject {
			return c.ItemPage(a, w)
		},
		"Shipping Address": func(c *Config, a fyne.App, w fyne.Window) fyne.CanvasObject {
			return c.ShippingAddressPage(a, w)
		},
		"Invoice": func(c *Config, a fyne.App, w fyne.Window) fyne.CanvasObject {
			return c.InvoicePage(a, w)
		},
	}

	sidebarKey = []string{
		"Company",
		"Payment Method",
		"Item",
		"Shipping Address",
		"Invoice",
	}
)

func (c *Config) sidebar(a fyne.App, w fyne.Window) *widget.List {
	list := widget.NewList(
		func() int { return len(sidebarKey) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Sidebar")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(sidebarKey[i])
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		key := sidebarKey[id]
		if fn, ok := sidebarData[key]; ok {
			slog.Info("Sidebar Action", "Clicked", key)
			con := fn(c, a, w)
			c.Cont.Trailing = con
			c.Cont.Refresh()
		}
	}

	return list
}
