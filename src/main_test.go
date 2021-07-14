package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

const Width = 2000
const Height = 1000

func TestGetStringWidth(t *testing.T) {
	width := 30
	snowString := getStringArray(width, DefaultSnowChar, DefaultFlakesRatio)
	runeCount := utf8.RuneCountInString(snowString)
	if width != runeCount {
		t.Error(fmt.Sprintf("Expcted runeCount %d, got %d", width, runeCount))
	}
}

func BenchmarkGetStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getStringConcat(Width, DefaultSnowChar, DefaultFlakesRatio)
	}
}

func BenchmarkGetStringArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getStringArray(Width, DefaultSnowChar, DefaultFlakesRatio)
	}
}

func BenchmarkGetStringAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getStringAppend(Width, DefaultSnowChar, DefaultFlakesRatio)
	}
}

func BenchmarkGetScreenConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getScreen(Height, Width, DefaultFlakesRatio, DefaultSnowChar, getStringConcat)
	}
}

func BenchmarkGetScreenAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getScreen(Height, Width, DefaultFlakesRatio, DefaultSnowChar, getStringAppend)
	}
}

func BenchmarkGetScreenArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getScreen(Height, Width, DefaultFlakesRatio, DefaultSnowChar, getStringArray)
	}
}
