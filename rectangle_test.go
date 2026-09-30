// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import "testing"

func TestRectangle_Location(t *testing.T) {
	r := Rectangle{X: 10, Y: 20, Width: 100, Height: 50}

	loc := r.Location()
	if loc.X != 10 || loc.Y != 20 {
		t.Errorf("Location() = %v; want Point{10, 20}", loc)
	}
}

func TestRectangle_SetLocation(t *testing.T) {
	r := Rectangle{X: 10, Y: 20, Width: 100, Height: 50}

	r.SetLocation(Point{X: 30, Y: 40})
	if r.X != 30 || r.Y != 40 {
		t.Errorf("SetLocation() failed: got %d, %d; want 30, 40", r.X, r.Y)
	}

	r2 := Rectangle{X: 10, Y: 20, Width: 100, Height: 50}
	r3 := r2.SetLocation(Point{X: 50, Y: 60})
	if r3.X != 50 || r3.Y != 60 || r3.Width != 100 || r3.Height != 50 {
		t.Errorf("SetLocation() returned %v; want Rectangle{50, 60, 100, 50}", r3)
	}
}
