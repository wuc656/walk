// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

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

	var te *walk.TextEdit

	if err := (declarative.MainWindow{
		Title:   "Walk Clipboard Example",
		MinSize: declarative.Size{Width: 300, Height: 200},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.PushButton{
				Text: "Copy",
				OnClicked: func() {
					if err := walk.Clipboard().SetText(te.Text()); err != nil {
						log.Print("Copy: ", err)
					}
				},
			},
			declarative.PushButton{
				Text: "Paste",
				OnClicked: func() {
					if text, err := walk.Clipboard().Text(); err != nil {
						log.Print("Paste: ", err)
					} else {
						te.SetText(text)
					}
				},
			},
			declarative.TextEdit{
				AssignTo: &te,
			},
		},
	}).Create(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}
