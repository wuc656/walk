// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package main

import (
	"log"
	"strings"

	"github.com/wuc656/walk"
	"github.com/wuc656/walk/declarative"
)

func main() {
	app, err := walk.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	var le *walk.LineEdit
	var wv *walk.WebView

	declarative.MainWindow{
		Icon:    declarative.Bind("'../img/' + icon(wv.URL) + '.ico'"),
		Title:   "Walk WebView Example'",
		MinSize: declarative.Size{Width: 800, Height: 600},
		Layout:  declarative.VBox{MarginsZero: true},
		Children: []declarative.Widget{
			declarative.LineEdit{
				AssignTo: &le,
				Text:     declarative.Bind("wv.URL"),
				OnKeyDown: func(key walk.Key) {
					if key == walk.KeyReturn {
						wv.SetURL(le.Text())
					}
				},
			},
			declarative.WebView{
				AssignTo: &wv,
				Name:     "wv",
				URL:      "https://github.com/wuc656/walk",
			},
		},
		Functions: map[string]func(args ...any) (any, error){
			"icon": func(args ...any) (any, error) {
				if strings.HasPrefix(args[0].(string), "https") {
					return "check", nil
				}

				return "stop", nil
			},
		},
	}.Create()

	app.Run()
}
