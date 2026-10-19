// Copyright 2016 The Walk Authors. All rights reserved.
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

	var slv, slh *walk.Slider
	var maxEdit, minEdit, valueEdit *walk.NumberEdit

	data := struct{ Min, Max, Value int }{0, 100, 30}

	declarative.MainWindow{
		Title:   "Walk Slider Example",
		MinSize: declarative.Size{Width: 320, Height: 240},
		Layout:  declarative.HBox{},
		Children: []declarative.Widget{
			declarative.Slider{
				AssignTo:    &slv,
				MinValue:    data.Min,
				MaxValue:    data.Max,
				Value:       data.Value,
				Orientation: declarative.Vertical,
				OnValueChanged: func() {
					data.Value = slv.Value()
					valueEdit.SetValue(float64(data.Value))

				},
			},
			declarative.Composite{
				Layout:        declarative.Grid{Columns: 3},
				StretchFactor: 4,
				Children: []declarative.Widget{
					declarative.Label{Text: "Min value"},
					declarative.Label{Text: "Value"},
					declarative.Label{Text: "Max value"},
					declarative.NumberEdit{
						AssignTo: &minEdit,
						Value:    float64(data.Min),
						OnValueChanged: func() {
							data.Min = int(minEdit.Value())
							slh.SetRange(data.Min, data.Max)
							slv.SetRange(data.Min, data.Max)
						},
					},
					declarative.NumberEdit{
						AssignTo: &valueEdit,
						Value:    float64(data.Value),
						OnValueChanged: func() {
							data.Value = int(valueEdit.Value())
							slh.SetValue(data.Value)
							slv.SetValue(data.Value)
						},
					},
					declarative.NumberEdit{
						AssignTo: &maxEdit,
						Value:    float64(data.Max),
						OnValueChanged: func() {
							data.Max = int(maxEdit.Value())
							slh.SetRange(data.Min, data.Max)
							slv.SetRange(data.Min, data.Max)
						},
					},
					declarative.Slider{
						ColumnSpan: 3,
						AssignTo:   &slh,
						MinValue:   data.Min,
						MaxValue:   data.Max,
						Value:      data.Value,
						OnValueChanged: func() {
							data.Value = slh.Value()
							valueEdit.SetValue(float64(data.Value))
						},
					},
					declarative.VSpacer{},
					declarative.PushButton{
						ColumnSpan: 3,
						Text:       "Print state",
						OnClicked: func() {
							log.Printf("H: < %d | %d | %d >\n", slh.MinValue(), slh.Value(), slh.MaxValue())
							log.Printf("V: < %d | %d | %d >\n", slv.MinValue(), slv.Value(), slv.MaxValue())
						},
					},
				},
			},
		},
	}.Create()

	app.Run()
}
