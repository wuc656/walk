// Copyright 2017 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

// This example demonstrates the status bar, including a size gripper
// attached to the bottom of the main window.
// The status bar has two items, one is dynamically updated and one includes an icon.
package main

import (
	"log"

	"github.com/wuc656/walk"
	"github.com/wuc656/walk/declarative"
)

func main() {
	app, err := walk.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	icon1, err := walk.NewIconFromFile("../img/check.ico")
	if err != nil {
		log.Fatal(err)
	}
	icon2, err := walk.NewIconFromFile("../img/stop.ico")
	if err != nil {
		log.Fatal(err)
	}

	var sbi *walk.StatusBarItem

	declarative.MainWindow{
		Title:   "Walk Statusbar Example",
		MinSize: declarative.Size{Width: 600, Height: 200},
		Layout:  declarative.VBox{MarginsZero: true},
		StatusBarItems: []declarative.StatusBarItem{
			declarative.StatusBarItem{
				AssignTo: &sbi,
				Icon:     icon1,
				Text:     "click",
				Width:    80,
				OnClicked: func() {
					if sbi.Text() == "click" {
						sbi.SetText("again")
						sbi.SetIcon(icon2)
					} else {
						sbi.SetText("click")
						sbi.SetIcon(icon1)
					}
				},
			},
			declarative.StatusBarItem{
				Text:        "left",
				ToolTipText: "no tooltip for me",
			},
			declarative.StatusBarItem{
				Text: "\tcenter",
			},
			declarative.StatusBarItem{
				Text: "\t\tright",
			},
			declarative.StatusBarItem{
				Icon:        icon1,
				ToolTipText: "An icon with a tooltip",
			},
		},
	}.Create()

	app.Run()
}
