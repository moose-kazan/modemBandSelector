package main

import (
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func actionHelpAbout() {
	dialog.NewInformation(
		"About",
		"Simple APP for Huawei HiLink modems.",
		w,
	).Show()
}

func actionWindowClose() {
	a.Quit()
}

func actionConnect() {
	modemIpEntry := widget.NewEntry()
	modemIpEntry.SetPlaceHolder("192.168.8.1")
	modemIpEntry.SetText("192.168.8.1")
	d := dialog.NewForm(
		"Connect to Modem/Router",
		"OK",
		"Cancel",
		[]*widget.FormItem{
			widget.NewFormItem(
				"Modem IP",
				modemIpEntry,
			),
		},
		func(b bool) {
			if !b {
				return
			}

			if modemIpEntry.Text != "" {
				err := h.Connect(modemIpEntry.Text)
				if err != nil {
					dialog.NewError(err, w).Show()
					return
				}
			}
		},
		w,
	)
	d.Show()
}
