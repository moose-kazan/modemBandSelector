package main

import (
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"fmt"
	"math"
)

func OnClose() {
	fmt.Println("Closed")
}

func GetData() {
	deviceInfo, err := h.DeviceInformation()
	if (err != nil) {
		dialog.NewError(err, w).Show()
		w.Close()
	}
	deviceInfoForm.Append("Device", widget.NewLabel(deviceInfo.DeviceName))
	deviceInfoForm.Append("S/N", widget.NewLabel(deviceInfo.SerialNumber))
	deviceInfoForm.Append("IMEI", widget.NewLabel(deviceInfo.Imei))
	deviceInfoForm.Append("IMSI", widget.NewLabel(deviceInfo.Imsi))
	deviceInfoForm.Append("HW Version", widget.NewLabel(deviceInfo.HardwareVersion))
	deviceInfoForm.Append("SW Version", widget.NewLabel(deviceInfo.SoftwareVersion))
	deviceInfoForm.Append("WebUI Version", widget.NewLabel(deviceInfo.WebUIVersion))
	deviceInfoForm.Append("Supported Modes", widget.NewLabel(deviceInfo.SupportMode))
	deviceInfoForm.Append("Work Mode", widget.NewLabel(deviceInfo.WorkMode))

	signalInfo, err := h.DeviceSignal()
	if (err != nil) {
		dialog.NewError(err, w).Show()
		w.Close()
	}
	signalInfoForm.Append("Band", widget.NewLabel(signalInfo.Band))
	signalInfoForm.Append("UL Bandwidth", widget.NewLabel(signalInfo.UlBandwidth))
	signalInfoForm.Append("DL Bandwidth", widget.NewLabel(signalInfo.DlBandwidth))

	netNetMode, err := h.NetNetMode()
	if (err != nil) {
		dialog.NewError(err, w).Show()
		w.Close()
	}
	bandList := make([]string, 0)
	lteband, _ := strconv.ParseInt(netNetMode.LTEBand, 16, 64)
	for i := 1; i < 66; i++ {
		tmp := int64(math.Pow(2, float64(i-1)))
		if (lteband & tmp) == tmp {
			bandList = append(bandList, fmt.Sprintf("%d", i))
		}
	}

	bandsListSelector := widget.NewCheckGroup(bandList, func(changed []string) {})
	selectBandsForm.Append("Bands", bandsListSelector)
	selectBandsForm.SubmitText = "Save"
	selectBandsForm.OnSubmit = func() {
		dialog.ShowInformation("Info", "Not implemented!", w)
	}
	selectBandsForm.Refresh()
}

func Connect() {
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
				w.Close()
			}

			if modemIpEntry.Text != "" {
				err := h.Connect(modemIpEntry.Text)
				if err != nil {
					dialog.NewError(err, w).Show()
					w.Close()
				}
				GetData()
			}
		},
		w,
	)
	d.Show()
}
