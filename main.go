package main

import (
	"modemBandSelector/internal/huaweiapi"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var (
	a fyne.App
	w fyne.Window
	h *huaweiapi.HuaweiAPI
)

func main() {
	h = huaweiapi.New()
	a = app.NewWithID("modemBandSelector")
	w = a.NewWindow("modemBandSelector")
	w.SetMainMenu(BuildMenu())
	w.Resize(fyne.NewSize(640, 480))
	h = huaweiapi.New()

	content := container.NewMax()
	w.SetContent(content)

	actionConnect()

	w.ShowAndRun()

	/*
		h.Connect("192.168.8.1")
		//fmt.Println(h.DeviceSignal())
		//fmt.Println(h.DeviceInformation())
		data, _ := h.NetNetMode()
		lteband, _ := strconv.ParseInt(data.LTEBand, 16, 64)
		for i := 1; i < 66; i++ {
			tmp := int64(math.Pow(2, float64(i-1)))
			if (lteband & tmp) == tmp {
				fmt.Println("Band:", i)
			}
		}
	*/
}
