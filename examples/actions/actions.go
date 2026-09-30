// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/wuc656/walk"
	"github.com/wuc656/walk/declarative"
)

var viewModes [4]*walk.Action
var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {
	app, err := walk.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	declarative.MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw := new(MyMainWindow)

	var openAction, showAboutBoxAction *walk.Action
	var recentMenu *walk.Menu
	var toggleSpecialModePB *walk.PushButton

	if err := (declarative.MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Walk Actions Example",
		MenuItems: []declarative.MenuItem{
			declarative.Menu{
				Text: "&File",
				Items: []declarative.MenuItem{
					declarative.Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						Image:       "../img/open.png",
						Enabled:     declarative.Bind("enabledCB.Checked"),
						Visible:     declarative.Bind("!openHiddenCB.Checked"),
						Shortcut:    declarative.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO},
						OnTriggered: mw.openAction_Triggered,
					},
					declarative.Menu{
						AssignTo: &recentMenu,
						Text:     "Recent",
					},
					declarative.Separator{},
					declarative.Action{
						Text:        "E&xit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			declarative.Menu{
				Text: "&View",
				Items: []declarative.MenuItem{
					declarative.Action{
						Text:    "Open / Special Enabled",
						Checked: declarative.Bind("enabledCB.Visible"),
					},
					declarative.Action{
						Text:    "Open Hidden",
						Checked: declarative.Bind("openHiddenCB.Visible"),
					},
				},
			},
			declarative.Menu{
				Text: "&Help",
				Items: []declarative.MenuItem{
					declarative.Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "About",
						OnTriggered: mw.showAboutBoxAction_Triggered,
					},
				},
			},
		},
		ToolBar: declarative.ToolBar{
			ButtonStyle: declarative.ToolBarButtonImageBeforeText,
			Items: []declarative.MenuItem{
				declarative.ActionRef{Action: &openAction},
				declarative.Menu{
					Text:  "New A",
					Image: "../img/document-new.png",
					Items: []declarative.MenuItem{
						declarative.Action{
							Text:        "A",
							OnTriggered: mw.newAction_Triggered,
						},
						declarative.Action{
							Text:        "B",
							OnTriggered: mw.newAction_Triggered,
						},
						declarative.Action{
							Text:        "C",
							OnTriggered: mw.newAction_Triggered,
						},
					},
					OnTriggered: mw.newAction_Triggered,
				},
				declarative.Separator{},
				declarative.Menu{
					Text:  "View",
					Image: "../img/document-properties.png",
					Items: []declarative.MenuItem{
						declarative.Action{
							AssignTo:    &viewModes[0],
							Text:        "X",
							OnTriggered: mw.changeViewAction_Triggered,
							Checkable:   true,
							Exclusive:   true,
						},
						declarative.Action{
							AssignTo:    &viewModes[1],
							Text:        "(Hidden)",
							OnTriggered: mw.changeViewAction_Triggered,
							Checkable:   true,
							Exclusive:   true,
							Visible:     false,
							Checked:     true,
						},
						declarative.Action{
							AssignTo:    &viewModes[2],
							Text:        "Y",
							OnTriggered: mw.changeViewAction_Triggered,
							Checkable:   true,
							Exclusive:   true,
						},
						declarative.Action{
							AssignTo:    &viewModes[3],
							Text:        "Z",
							OnTriggered: mw.changeViewAction_Triggered,
							Checkable:   true,
							Exclusive:   true,
						},
					},
				},
				declarative.Separator{},
				declarative.Action{
					Text:        "Special",
					Image:       "../img/system-shutdown.png",
					Enabled:     declarative.Bind("isSpecialMode && enabledCB.Checked"),
					OnTriggered: mw.specialAction_Triggered,
				},
			},
		},
		ContextMenuItems: []declarative.MenuItem{
			declarative.ActionRef{Action: &showAboutBoxAction},
		},
		MinSize: declarative.Size{Width: 300, Height: 200},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.CheckBox{
				Name:    "enabledCB",
				Text:    "Open / Special Enabled",
				Checked: true,
				Accessibility: declarative.Accessibility{
					Help: "Enables Open and Special",
				},
			},
			declarative.CheckBox{
				Name:    "openHiddenCB",
				Text:    "Open Hidden",
				Checked: true,
			},
			declarative.PushButton{
				AssignTo: &toggleSpecialModePB,
				Text:     "Enable Special Mode",
				OnClicked: func() {
					isSpecialMode.SetSatisfied(!isSpecialMode.Satisfied())

					if isSpecialMode.Satisfied() {
						toggleSpecialModePB.SetText("Disable Special Mode")
					} else {
						toggleSpecialModePB.SetText("Enable Special Mode")
					}
				},
				Accessibility: declarative.Accessibility{
					Help: "Toggles special mode",
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	addRecentFileActions := func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(mw.openAction_Triggered)
			recentMenu.Actions().Add(a)
		}
	}

	addRecentFileActions("Foo", "Bar", "Baz")

	app.Run()
}

func (mw *MyMainWindow) openAction_Triggered() {
	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) newAction_Triggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) changeViewAction_Triggered() {
	var msg strings.Builder
	msg.WriteString("Current view mode:\n")
	for _, m := range viewModes {
		if m.Checked() {
			fmt.Fprintf(&msg, " - %s\n", m.Text())
		}
	}
	walk.MsgBox(mw, "Change View", msg.String(), walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
