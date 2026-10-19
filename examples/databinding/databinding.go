// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/wuc656/walk"

	"github.com/wuc656/walk/declarative"
)

func main() {
	walk.AppendToWalkInit(func() {
		walk.FocusEffect, _ = walk.NewBorderGlowEffect(walk.RGB(0, 63, 255))
		walk.InteractionEffect, _ = walk.NewDropShadowEffect(walk.RGB(63, 63, 63))
		walk.ValidationErrorEffect, _ = walk.NewBorderGlowEffect(walk.RGB(255, 0, 0))
	})

	app, err := walk.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	var mw *walk.MainWindow
	var outTE *walk.TextEdit

	animal := new(Animal)

	if err := (declarative.MainWindow{
		AssignTo: &mw,
		Title:    "Walk Data Binding Example",
		MinSize:  declarative.Size{Width: 300, Height: 200},
		Layout:   declarative.VBox{},
		Children: []declarative.Widget{
			declarative.PushButton{
				Text: "Edit Animal",
				OnClicked: func() {
					if cmd, err := RunAnimalDialog(mw, animal); err != nil {
						log.Print(err)
					} else if cmd == walk.DlgCmdOK {
						outTE.SetText(fmt.Sprintf("%+v", animal))
					}
				},
			},
			declarative.Label{
				Text: "animal:",
			},
			declarative.TextEdit{
				AssignTo: &outTE,
				ReadOnly: true,
				Text:     fmt.Sprintf("%+v", animal),
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	app.Run()
}

type Animal struct {
	Name          string
	ArrivalDate   time.Time
	SpeciesId     int
	Speed         int
	Sex           Sex
	Weight        float64
	PreferredFood string
	Domesticated  bool
	Remarks       string
	Patience      time.Duration
}

func (a *Animal) PatienceField() *DurationField {
	return &DurationField{&a.Patience}
}

type Species struct {
	Id   int
	Name string
}

func KnownSpecies() []*Species {
	return []*Species{
		{1, "Dog"},
		{2, "Cat"},
		{3, "Bird"},
		{4, "Fish"},
		{5, "Elephant"},
	}
}

type DurationField struct {
	p *time.Duration
}

func (*DurationField) CanSet() bool { return true }
func (f *DurationField) Get() any   { return f.p.String() }
func (f *DurationField) Set(v any) error {
	x, err := time.ParseDuration(v.(string))
	if err == nil {
		*f.p = x
	}
	return err
}
func (f *DurationField) Zero() any { return "" }

type Sex byte

const (
	SexMale Sex = 1 + iota
	SexFemale
	SexHermaphrodite
)

func RunAnimalDialog(owner walk.Form, animal *Animal) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	if err := (declarative.Dialog{
		AssignTo:      &dlg,
		Title:         declarative.Bind("'Animal Details' + (animal.Name == '' ? '' : ' - ' + animal.Name)"),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: declarative.DataBinder{
			AssignTo:       &db,
			Name:           "animal",
			DataSource:     animal,
			ErrorPresenter: declarative.ToolTipErrorPresenter{},
		},
		MinSize: declarative.Size{Width: 300, Height: 300},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Composite{
				Layout: declarative.Grid{Columns: 2},
				Children: []declarative.Widget{
					declarative.Label{
						Text: "Name:",
					},
					declarative.LineEdit{
						Text: declarative.Bind("Name"),
					},

					declarative.Label{
						Text: "Arrival Date:",
					},
					declarative.DateEdit{
						Date: declarative.Bind("ArrivalDate"),
					},

					declarative.Label{
						Text: "Species:",
					},
					declarative.ComboBox{
						Value:         declarative.Bind("SpeciesId", declarative.SelRequired{}),
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         KnownSpecies(),
					},

					declarative.Label{
						Text: "Speed:",
					},
					declarative.Slider{
						Value: declarative.Bind("Speed"),
					},

					declarative.RadioButtonGroupBox{
						ColumnSpan: 2,
						Title:      "Sex",
						Layout:     declarative.HBox{},
						DataMember: "Sex",
						Buttons: []declarative.RadioButton{
							{Text: "Male", Value: SexMale},
							{Text: "Female", Value: SexFemale},
							{Text: "Hermaphrodite", Value: SexHermaphrodite},
						},
					},

					declarative.Label{
						Text: "Weight:",
					},
					declarative.NumberEdit{
						Value:    declarative.Bind("Weight", declarative.Range{Min: 0.01, Max: 9999.99}),
						Suffix:   " kg",
						Decimals: 2,
					},

					declarative.Label{
						Text: "Preferred Food:",
					},
					declarative.ComboBox{
						Editable: true,
						Value:    declarative.Bind("PreferredFood"),
						Model:    []string{"Fruit", "Grass", "Fish", "Meat"},
					},

					declarative.Label{
						Text: "Domesticated:",
					},
					declarative.CheckBox{
						Checked: declarative.Bind("Domesticated"),
					},

					declarative.VSpacer{
						ColumnSpan: 2,
						Size:       8,
					},

					declarative.Label{
						ColumnSpan: 2,
						Text:       "Remarks:",
					},
					declarative.TextEdit{
						ColumnSpan: 2,
						MinSize:    declarative.Size{Width: 100, Height: 50},
						Text:       declarative.Bind("Remarks"),
					},

					declarative.Label{
						ColumnSpan: 2,
						Text:       "Patience:",
					},
					declarative.LineEdit{
						ColumnSpan: 2,
						Text:       declarative.Bind("PatienceField"),
					},
				},
			},
			declarative.Composite{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.HSpacer{},
					declarative.PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
						},
					},
					declarative.PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}).Create(owner); err != nil {
		return 0, err
	}

	return dlg.Run(), nil
}
