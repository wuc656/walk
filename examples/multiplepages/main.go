// Copyright 2017 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package main

import (
	"bytes"
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

	mw := new(AppMainWindow)

	cfg := &MultiPageMainWindowConfig{
		Name:    "mainWindow",
		MinSize: declarative.Size{Width: 600, Height: 400},
		MenuItems: []declarative.MenuItem{
			declarative.Menu{
				Text: "&Help",
				Items: []declarative.MenuItem{
					declarative.Action{
						Text:        "About",
						OnTriggered: func() { mw.aboutAction_Triggered() },
					},
				},
			},
		},
		OnCurrentPageChanged: func() {
			mw.updateTitle(mw.CurrentPageTitle())
		},
		PageCfgs: []PageConfig{
			{"Foo", "document-new.png", newFooPage},
			{"Bar", "document-properties.png", newBarPage},
			{"Baz", "system-shutdown.png", newBazPage},
		},
	}

	mpmw, err := NewMultiPageMainWindow(cfg)
	if err != nil {
		panic(err)
	}

	mw.MultiPageMainWindow = mpmw

	mw.updateTitle(mw.CurrentPageTitle())

	app.Run()
}

type AppMainWindow struct {
	*MultiPageMainWindow
}

func (mw *AppMainWindow) updateTitle(prefix string) {
	var buf bytes.Buffer

	if prefix != "" {
		buf.WriteString(prefix)
		buf.WriteString(" - ")
	}

	buf.WriteString("Walk Multiple Pages Example")

	mw.SetTitle(buf.String())
}

func (mw *AppMainWindow) aboutAction_Triggered() {
	walk.MsgBox(mw,
		"About Walk Multiple Pages Example",
		"An example that demonstrates a main window that supports multiple pages.",
		walk.MsgBoxOK|walk.MsgBoxIconInformation)
}

type FooPage struct {
	*walk.Composite
}

func newFooPage(parent walk.Container) (Page, error) {
	p := new(FooPage)

	if err := (declarative.Composite{
		AssignTo: &p.Composite,
		Name:     "fooPage",
		Layout:   declarative.HBox{},
		Children: []declarative.Widget{
			declarative.HSpacer{},
			declarative.Label{Text: "I'm the Foo page"},
			declarative.HSpacer{},
		},
	}).Create(declarative.NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}

type BarPage struct {
	*walk.Composite
}

func newBarPage(parent walk.Container) (Page, error) {
	p := new(BarPage)

	if err := (declarative.Composite{
		AssignTo: &p.Composite,
		Name:     "barPage",
		Layout:   declarative.HBox{},
		Children: []declarative.Widget{
			declarative.HSpacer{},
			declarative.Label{Text: "I'm the Bar page"},
			declarative.HSpacer{},
		},
	}).Create(declarative.NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}

type BazPage struct {
	*walk.Composite
}

func newBazPage(parent walk.Container) (Page, error) {
	p := new(BazPage)

	if err := (declarative.Composite{
		AssignTo: &p.Composite,
		Name:     "bazPage",
		Layout:   declarative.HBox{},
		Children: []declarative.Widget{
			declarative.HSpacer{},
			declarative.Label{Text: "I'm the Baz page"},
			declarative.HSpacer{},
		},
	}).Create(declarative.NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}
