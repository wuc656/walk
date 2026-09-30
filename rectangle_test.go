// Copyright 2024 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
)

func TestRectangle_SetLocation(t *testing.T) {
	rect := Rectangle{
		X:      10,
		Y:      20,
		Width:  100,
		Height: 200,
	}

	newPoint := Point{X: 50, Y: 60}

	expectedRect := Rectangle{
		X:      50,
		Y:      60,
		Width:  100,
		Height: 200,
	}

	updatedRect := rect.SetLocation(newPoint)

	if rect != expectedRect {
		t.Errorf("SetLocation modified original receiver to %#v, want %#v", rect, expectedRect)
	}

	if updatedRect != expectedRect {
		t.Errorf("SetLocation returned %#v, want %#v", updatedRect, expectedRect)
	}
}
