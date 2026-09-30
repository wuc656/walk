// Copyright 2023 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
)

func TestFormatFloatGrouped(t *testing.T) {
	oldGroupSepS := groupSepS
	oldDecimalSepS := decimalSepS

	groupSepS = ","
	decimalSepS = "."

	defer func() {
		groupSepS = oldGroupSepS
		decimalSepS = oldDecimalSepS
	}()

	cases := []struct {
		f    float64
		prec int
		want string
	}{
		{1234.56, 0, "1,234"},
		{1234.96, 0, "1,235"},
		{1234.56, 2, "1,234.56"},
		{1234.56, 4, "1,234.5600"},
		{1234567.89, 2, "1,234,567.89"},
		{-1234.56, 0, "-1,234"},
		{-1234.96, 0, "-1,235"},
		{-1234.56, 2, "-1,234.56"},
		{-1234.56, 4, "-1,234.5600"},
		{-1234567.89, 2, "-1,234,567.89"},
		{0.12, 2, "0.12"},
		{-0.12, 2, "-0.12"},
		{123.45, 2, "123.45"},
		{1000.0, 2, "1,000.00"},
		{12.34, 2, "12.34"},
		{-12.34, 2, "-12.34"},
		{0.0, 2, "0.00"},
	}

	for _, tc := range cases {
		got := FormatFloatGrouped(tc.f, tc.prec)
		if got != tc.want {
			t.Errorf("FormatFloatGrouped(%v, %v) = %q; want %q", tc.f, tc.prec, got, tc.want)
		}
	}
}
