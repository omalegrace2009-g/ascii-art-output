package main

import (
	"reflect"
	"testing"
)

func testBanner() map[rune][]string {
	return map[rune][]string{
		'A': {
			"A1", "A2", "A3", "A4",
			"A5", "A6", "A7", "A8",
		},
		'B': {
			"B1", "B2", "B3", "B4",
			"B5", "B6", "B7", "B8",
		},
	}
}
func TestRender_EmptyInput(t *testing.T) {
	got := Render("", testBanner())

	if len(got) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(got))
	}

	for i, line := range got {
		if line != "" {
			t.Errorf("line %d expected empty string, got %q", i, line)
		}
	}
}
func TestRender_SingleCharacter(t *testing.T) {
	want := []string{
		"A1", "A2", "A3", "A4",
		"A5", "A6", "A7", "A8",
	}

	got := Render("A", testBanner())

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
func TestRender_MultipleCharacters(t *testing.T) {
	want := []string{
		"A1B1",
		"A2B2",
		"A3B3",
		"A4B4",
		"A5B5",
		"A6B6",
		"A7B7",
		"A8B8",
	}

	got := Render("AB", testBanner())

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
func TestRender_RepeatedCharacters(t *testing.T) {
	want := []string{
		"A1A1",
		"A2A2",
		"A3A3",
		"A4A4",
		"A5A5",
		"A6A6",
		"A7A7",
		"A8A8",
	}

	got := Render("AA", testBanner())

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
func TestRender_UnknownCharacter(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for unknown character")
		}
	}()

	Render("@", testBanner())
}
func TestRender_InvalidBanner(t *testing.T) {
	banner := map[rune][]string{
		'A': {"A1"},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for invalid banner")
		}
	}()

	Render("A", banner)
}
func TestRender_LineCount(t *testing.T) {
	got := Render("AB", testBanner())

	if len(got) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(got))
	}
}
