//go:build windows

package walk

import (
	"strconv"
	"testing"
	"golang.org/x/sys/windows"
)

func genMockButtons(count int, noteType string) []TaskDialogCustomButton {
	var btns []TaskDialogCustomButton
	for i := 0; i < count; i++ {
		if noteType == "short_ascii" {
			btns = append(btns, TaskDialogCustomButton{MainText: "Button", Note: "Note"})
		} else if noteType == "long_ascii" {
			btns = append(btns, TaskDialogCustomButton{MainText: "This is a very long button text", Note: "And this is a very long note for the button"})
		} else if noteType == "unicode" {
			btns = append(btns, TaskDialogCustomButton{MainText: "按鈕", Note: "這是一個備註"})
		}
	}
	return btns
}

func BenchmarkTaskDialogCustomButtons(b *testing.B) {
	for _, count := range []int{1, 2, 5, 10} {
		for _, typ := range []string{"short_ascii", "long_ascii", "unicode"} {
			btns := genMockButtons(count, typ)
			b.Run("Count="+strconv.Itoa(count)+"_Type="+typ, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, btn := range btns {
						var text string
						if btn.Note != "" {
							text = btn.MainText + "\n" + btn.Note
						} else {
							text = btn.MainText
						}
						_, _ = windows.UTF16PtrFromString(text)
					}
				}
			})
		}
	}
}
