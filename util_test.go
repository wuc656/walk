// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"math"
	"testing"
)

func TestParseFloat(t *testing.T) {
	origDecimalSepS := decimalSepS
	origGroupSepS := groupSepS
	defer func() {
		decimalSepS = origDecimalSepS
		groupSepS = origGroupSepS
	}()

	tests := []struct {
		name        string
		input       string
		decimalSep  string
		groupSep    string
		want        float64
		wantErr     bool
	}{
		{
			name:       "US locale standard",
			input:      "1,234.56",
			decimalSep: ".",
			groupSep:   ",",
			want:       1234.56,
			wantErr:    false,
		},
		{
			name:       "German locale standard",
			input:      "1.234,56",
			decimalSep: ",",
			groupSep:   ".",
			want:       1234.56,
			wantErr:    false,
		},
		{
			name:       "French locale space",
			input:      "1 234,56",
			decimalSep: ",",
			groupSep:   " ",
			want:       1234.56,
			wantErr:    false,
		},
		{
			name:       "No group separator",
			input:      "1234.56",
			decimalSep: ".",
			groupSep:   ",",
			want:       1234.56,
			wantErr:    false,
		},
		{
			name:       "Negative US locale",
			input:      "-1,234.56",
			decimalSep: ".",
			groupSep:   ",",
			want:       -1234.56,
			wantErr:    false,
		},
		{
			name:       "Trims space",
			input:      "  1,234.56  ",
			decimalSep: ".",
			groupSep:   ",",
			want:       1234.56,
			wantErr:    false,
		},
		{
			name:       "Invalid input",
			input:      "abc",
			decimalSep: ".",
			groupSep:   ",",
			want:       0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decimalSepS = tt.decimalSep
			groupSepS = tt.groupSep

			got, err := ParseFloat(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("ParseFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
