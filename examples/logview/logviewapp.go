// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package main

import (
	"log"
	"time"

	"github.com/wuc656/walk"

	"github.com/wuc656/walk/declarative"
)

func main() {
	app, err := walk.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	var mw *walk.MainWindow

	if err := (declarative.MainWindow{
		AssignTo: &mw,
		Title:    "Walk LogView Example",
		MinSize:  declarative.Size{Width: 320, Height: 240},
		Size:     declarative.Size{Width: 400, Height: 600},
		Layout:   declarative.VBox{MarginsZero: true},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	lv, err := NewLogView(mw)
	if err != nil {
		log.Fatal(err)
	}

	lv.PostAppendText("XXX")
	log.SetOutput(lv)

	go func() {
		for range 10000 {
			time.Sleep(100 * time.Millisecond)
			log.Println("Text")
		}
	}()

	app.Run()
}
