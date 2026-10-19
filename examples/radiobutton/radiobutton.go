// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package main

import (
	"fmt"
	"log"

	"github.com/wuc656/walk"

	"github.com/wuc656/walk/declarative"
)

type Foo struct {
	Bar string
	Baz int
}

func main() {
	app, err := walk.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	foo := &Foo{"b", 0}

	declarative.MainWindow{
		Title:   "Walk RadioButton Example",
		MinSize: declarative.Size{Width: 320, Height: 240},
		Layout:  declarative.VBox{},
		DataBinder: declarative.DataBinder{
			DataSource: foo,
			AutoSubmit: true,
			OnSubmitted: func() {
				fmt.Println(foo)
			},
		},
		Children: []declarative.Widget{
			// RadioButtonGroup is needed for data binding only.
			declarative.RadioButtonGroup{
				DataMember: "Bar",
				Buttons: []declarative.RadioButton{
					declarative.RadioButton{
						Name:  "aRB",
						Text:  "A",
						Value: "a",
					},
					declarative.RadioButton{
						Name:  "bRB",
						Text:  "B",
						Value: "b",
					},
					declarative.RadioButton{
						Name:  "cRB",
						Text:  "C",
						Value: "c",
					},
				},
			},
			declarative.Label{
				Text:    "A",
				Enabled: declarative.Bind("aRB.Checked"),
			},
			declarative.Label{
				Text:    "B",
				Enabled: declarative.Bind("bRB.Checked"),
			},
			declarative.Label{
				Text:    "C",
				Enabled: declarative.Bind("cRB.Checked"),
			},
			declarative.RadioButtonGroup{
				DataMember: "Baz",
				Buttons: []declarative.RadioButton{
					declarative.RadioButton{
						Name:  "oneRB",
						Text:  "1",
						Value: 1,
					},
					declarative.RadioButton{
						Name:  "twoRB",
						Text:  "2",
						Value: 2,
					},
					declarative.RadioButton{
						Name:  "threeRB",
						Text:  "3",
						Value: 3,
					},
				},
			},
			declarative.Label{
				Text:    "1",
				Enabled: declarative.Bind("oneRB.Checked"),
			},
			declarative.Label{
				Text:    "2",
				Enabled: declarative.Bind("twoRB.Checked"),
			},
			declarative.Label{
				Text:    "3",
				Enabled: declarative.Bind("threeRB.Checked"),
			},
		},
	}.Create()

	app.Run()
}
