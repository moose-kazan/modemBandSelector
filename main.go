package main

import (
	"modemBandSelector/internal/huaweiapi"

	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	a fyne.App
	w fyne.Window
	deviceInfoForm *widget.Form
	signalInfoForm *widget.Form
	selectBandsForm *widget.Form
	h *huaweiapi.HuaweiAPI
)

func main() {
	os.Setenv("FYNE_THEME", "light")
	h = huaweiapi.New()
	a = app.NewWithID("modemBandSelector")
	w = a.NewWindow("modemBandSelector")
	w.SetOnClosed(OnClose)
	deviceInfoForm = widget.NewForm()
	signalInfoForm = widget.NewForm()
	selectBandsForm = widget.NewForm()

	w.Resize(fyne.NewSize(640, 480))
	h = huaweiapi.New()

	contentInfo := container.NewVBox(
		widget.NewLabelWithStyle("Device Info", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		deviceInfoForm,
		widget.NewLabelWithStyle("Signal Info", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		signalInfoForm,
	)
	contentBands := container.NewVBox(
		widget.NewLabelWithStyle("Supported Bands", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		selectBandsForm,
	)
	content := container.NewHBox(
		contentInfo,
		contentBands,
	)
	w.SetContent(content)

	Connect()
	
	w.ShowAndRun()
}
