package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestGetStringWidth(t *testing.T) {
	width := 30
	snowString := getString(width, DefaultSnowChar, DefaultFlakesRatio)
	runeCount := utf8.RuneCountInString(snowString)
	if width != runeCount {
		t.Error(fmt.Sprintf("Expcted runeCount %d, got %d", width, runeCount))
	}
}

func BenchmarkGetString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getString(80, DefaultSnowChar, DefaultFlakesRatio)
	}
}

func BenchmarkGetStringWidth1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getString(1000, DefaultSnowChar, DefaultFlakesRatio)
	}
}

func BenchmarkGetStringAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getStringAppend(80, DefaultSnowChar, DefaultFlakesRatio)
	}
}

func BenchmarkGetScreen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getScreen(40, 80, DefaultFlakesRatio, DefaultSnowChar)
	}
}

func BenchmarkGetScreenAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getScreenAppend(40, 80, DefaultFlakesRatio, DefaultSnowChar)
	}
}
