// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"math"
	"testing"
)

func TestFormatFloat(t *testing.T) {
	origDecimalSepS := decimalSepS
	origGroupSepS := groupSepS
	decimalSepS = "."
	groupSepS = ","
	defer func() {
		decimalSepS = origDecimalSepS
		groupSepS = origGroupSepS
	}()

	testCases := []struct {
		f       float64
		prec    int
		want    string
		grouped bool
	}{
		// FormatFloat cases
		{0.0, 2, "0.00", false},
		{1234.5, 2, "1234.50", false},
		{1234.5, 0, "1234", false},
		{-1234.5, 2, "-1234.50", false},
		{math.NaN(), 2, "NaN", false},
		{math.Inf(1), 2, "+Inf", false},
		{math.Inf(-1), 2, "-Inf", false},

		// FormatFloatGrouped cases
		{0.0, 2, "0.00", true},
		{1234.5, 2, "1,234.50", true},
		{1234.5, 0, "1,234", true},
		{1234567.89, 2, "1,234,567.89", true},
		{-1234.5, 2, "-1,234.50", true},
		{-1234567.89, 2, "-1,234,567.89", true},
		{math.NaN(), 2, "NaN", true},
		{math.Inf(1), 2, "+Inf", true},
		{math.Inf(-1), 2, "-Inf", true},
	}

	for _, tc := range testCases {
		var got string
		if tc.grouped {
			got = FormatFloatGrouped(tc.f, tc.prec)
		} else {
			got = FormatFloat(tc.f, tc.prec)
		}
		if got != tc.want {
			t.Errorf("FormatFloat(f=%v, prec=%v, grouped=%v) = %q; want %q", tc.f, tc.prec, tc.grouped, got, tc.want)
		}
	}
}
