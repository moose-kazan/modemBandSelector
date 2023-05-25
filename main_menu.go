package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func newMenuItem(label string, action func(), Icon fyne.Resource, Shortcut fyne.Shortcut) *fyne.MenuItem {
	m := fyne.NewMenuItem(label, action)
	m.Icon = Icon
	m.Shortcut = Shortcut
	return m
}

func BuildMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu(
			"Connection",
			newMenuItem("Connect", actionConnect, theme.DocumentIcon(), nil),
			newMenuItem("Quit", actionWindowClose, theme.LogoutIcon(), nil),
		),
		fyne.NewMenu(
			"Help",
			newMenuItem("About", actionHelpAbout, theme.InfoIcon(), nil),
		),
	)
}
