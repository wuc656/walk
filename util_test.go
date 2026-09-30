// Copyright 2024 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"reflect"
	"testing"

	"github.com/wuc656/win"
)

func TestMARGINSFrom96DPI(t *testing.T) {
	tests := []struct {
		name  string
		value win.MARGINS
		dpi   int
		want  win.MARGINS
	}{
		{
			name: "96 DPI (100% scale)",
			value: win.MARGINS{
				LeftWidth:    10,
				RightWidth:   20,
				TopHeight:    30,
				BottomHeight: 40,
			},
			dpi: 96,
			want: win.MARGINS{
				LeftWidth:    10,
				RightWidth:   20,
				TopHeight:    30,
				BottomHeight: 40,
			},
		},
		{
			name: "144 DPI (150% scale)",
			value: win.MARGINS{
				LeftWidth:    10,
				RightWidth:   20,
				TopHeight:    30,
				BottomHeight: 40,
			},
			dpi: 144,
			want: win.MARGINS{
				LeftWidth:    15,
				RightWidth:   30,
				TopHeight:    45,
				BottomHeight: 60,
			},
		},
		{
			name: "192 DPI (200% scale)",
			value: win.MARGINS{
				LeftWidth:    10,
				RightWidth:   20,
				TopHeight:    30,
				BottomHeight: 40,
			},
			dpi: 192,
			want: win.MARGINS{
				LeftWidth:    20,
				RightWidth:   40,
				TopHeight:    60,
				BottomHeight: 80,
			},
		},
		{
			name: "Zero value",
			value: win.MARGINS{
				LeftWidth:    0,
				RightWidth:   0,
				TopHeight:    0,
				BottomHeight: 0,
			},
			dpi: 144,
			want: win.MARGINS{
				LeftWidth:    0,
				RightWidth:   0,
				TopHeight:    0,
				BottomHeight: 0,
			},
		},
		{
			name: "Rounding check",
			value: win.MARGINS{
				LeftWidth:    3, // 3 * 1.5 = 4.5 -> 5
				RightWidth:   5, // 5 * 1.5 = 7.5 -> 8
				TopHeight:    7, // 7 * 1.5 = 10.5 -> 11
				BottomHeight: 9, // 9 * 1.5 = 13.5 -> 14
			},
			dpi: 144,
			want: win.MARGINS{
				LeftWidth:    5,
				RightWidth:   8,
				TopHeight:    11,
				BottomHeight: 14,
			},
		},
		{
			name: "Negative values",
			value: win.MARGINS{
				LeftWidth:    -10,
				RightWidth:   -20,
				TopHeight:    -30,
				BottomHeight: -40,
			},
			dpi: 144,
			want: win.MARGINS{
				LeftWidth:    -15,
				RightWidth:   -30,
				TopHeight:    -45,
				BottomHeight: -60,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MARGINSFrom96DPI(tt.value, tt.dpi)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MARGINSFrom96DPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
