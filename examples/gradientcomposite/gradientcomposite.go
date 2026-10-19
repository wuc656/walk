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

	declarative.MainWindow{
		Title:   "Walk GradientComposite Example",
		MinSize: declarative.Size{Width: 400, Height: 0},
		Background: declarative.GradientBrush{
			Vertexes: []walk.GradientVertex{
				{X: 0, Y: 0, Color: walk.RGB(255, 255, 127)},
				{X: 1, Y: 0, Color: walk.RGB(127, 191, 255)},
				{X: 0.5, Y: 0.5, Color: walk.RGB(255, 255, 255)},
				{X: 1, Y: 1, Color: walk.RGB(127, 255, 127)},
				{X: 0, Y: 1, Color: walk.RGB(255, 127, 127)},
			},
			Triangles: []walk.GradientTriangle{
				{Vertex1: 0, Vertex2: 1, Vertex3: 2},
				{Vertex1: 1, Vertex2: 3, Vertex3: 2},
				{Vertex1: 3, Vertex2: 4, Vertex3: 2},
				{Vertex1: 4, Vertex2: 0, Vertex3: 2},
			},
		},
		Layout: declarative.HBox{Margins: declarative.Margins{Left: 100, Top: 100, Right: 100, Bottom: 100}},
		Children: []declarative.Widget{
			declarative.GradientComposite{
				Border:   true,
				Vertical: declarative.Bind("verticalCB.Checked"),
				Color1:   declarative.Bind("rgb(c1RedSld.Value, c1GreenSld.Value, c1BlueSld.Value)"),
				Color2:   declarative.Bind("rgb(c2RedSld.Value, c2GreenSld.Value, c2BlueSld.Value)"),
				Layout:   declarative.HBox{},
				Children: []declarative.Widget{
					declarative.GroupBox{
						Title:  "Gradient Parameters",
						Layout: declarative.VBox{},
						Children: []declarative.Widget{
							declarative.CheckBox{Name: "verticalCB", Text: "Vertical", Checked: true},
							declarative.GroupBox{
								Title:  "Color1",
								Layout: declarative.Grid{Columns: 2},
								Children: []declarative.Widget{
									declarative.Label{Text: "Red:"},
									declarative.Slider{Name: "c1RedSld", Tracking: true, MaxValue: 255, Value: 95},
									declarative.Label{Text: "Green:"},
									declarative.Slider{Name: "c1GreenSld", Tracking: true, MaxValue: 255, Value: 191},
									declarative.Label{Text: "Blue:"},
									declarative.Slider{Name: "c1BlueSld", Tracking: true, MaxValue: 255, Value: 255},
								},
							},
							declarative.GroupBox{
								Title:  "Color2",
								Layout: declarative.Grid{Columns: 2},
								Children: []declarative.Widget{
									declarative.Label{Text: "Red:"},
									declarative.Slider{Name: "c2RedSld", Tracking: true, MaxValue: 255, Value: 239},
									declarative.Label{Text: "Green:"},
									declarative.Slider{Name: "c2GreenSld", Tracking: true, MaxValue: 255, Value: 63},
									declarative.Label{Text: "Blue:"},
									declarative.Slider{Name: "c2BlueSld", Tracking: true, MaxValue: 255, Value: 0},
								},
							},
						},
					},
				},
			},
		},
		Functions: map[string]func(args ...any) (any, error){
			"rgb": func(args ...any) (any, error) {
				return walk.RGB(byte(args[0].(float64)), byte(args[1].(float64)), byte(args[2].(float64))), nil
			},
		},
	}.Create()

	app.Run()
}
