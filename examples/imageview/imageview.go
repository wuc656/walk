// Copyright 2017 The Walk Authors. All rights reserved.
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

	walk.Resources.SetRootDirPath("../img")

	type Mode struct {
		Name  string
		Value declarative.ImageViewMode
	}

	modes := []Mode{
		{"ImageViewModeIdeal", declarative.ImageViewModeIdeal},
		{"ImageViewModeCorner", declarative.ImageViewModeCorner},
		{"ImageViewModeCenter", declarative.ImageViewModeCenter},
		{"ImageViewModeShrink", declarative.ImageViewModeShrink},
		{"ImageViewModeZoom", declarative.ImageViewModeZoom},
		{"ImageViewModeStretch", declarative.ImageViewModeStretch},
	}

	var widgets []declarative.Widget

	for _, mode := range modes {
		widgets = append(widgets,
			declarative.Label{
				Text: mode.Name,
			},
			declarative.ImageView{
				Background: declarative.SolidColorBrush{Color: walk.RGB(255, 191, 0)},
				Image:      "open.png",
				Margin:     10,
				Mode:       mode.Value,
			},
		)
	}

	declarative.MainWindow{
		Title:    "Walk ImageView Example",
		Size:     declarative.Size{Width: 400, Height: 600},
		Layout:   declarative.Grid{Columns: 2},
		Children: widgets,
	}.Create()

	app.Run()
}
