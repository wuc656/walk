// Copyright 2023 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
)

func TestRectangle_SetSize(t *testing.T) {
	rect := Rectangle{X: 10, Y: 20, Width: 30, Height: 40}
	newSize := Size{Width: 50, Height: 60}

	returnedRect := rect.SetSize(newSize)

	// Verify original rectangle is updated
	if rect.Width != newSize.Width || rect.Height != newSize.Height {
		t.Errorf("Expected original rectangle dimensions to be updated to %v, got Size{%d, %d}", newSize, rect.Width, rect.Height)
	}
	if rect.X != 10 || rect.Y != 20 {
		t.Errorf("Expected original rectangle coordinates to remain {10, 20}, got {%d, %d}", rect.X, rect.Y)
	}

	// Verify returned rectangle has correct values
	if returnedRect.Width != newSize.Width || returnedRect.Height != newSize.Height {
		t.Errorf("Expected returned rectangle dimensions to be %v, got Size{%d, %d}", newSize, returnedRect.Width, returnedRect.Height)
	}
	if returnedRect.X != 10 || returnedRect.Y != 20 {
		t.Errorf("Expected returned rectangle coordinates to be {10, 20}, got {%d, %d}", returnedRect.X, returnedRect.Y)
	}
}
