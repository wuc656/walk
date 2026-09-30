// Copyright 2026 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
)

func TestRectangle_Left(t *testing.T) {
	testCases := []struct {
		name string
		rect Rectangle
		want int
	}{
		{
			name: "Positive X",
			rect: Rectangle{X: 10, Y: 20, Width: 100, Height: 50},
			want: 10,
		},
		{
			name: "Zero X",
			rect: Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
			want: 0,
		},
		{
			name: "Negative X",
			rect: Rectangle{X: -5, Y: 15, Width: 200, Height: 100},
			want: -5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.rect.Left()
			if got != tc.want {
				t.Errorf("Rectangle.Left() = %v, want %v", got, tc.want)
			}
		})
	}
}
