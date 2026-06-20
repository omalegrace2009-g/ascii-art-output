package main

import (
	"strings"
	"testing"
)

func mockBanner() map[rune][]string {
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
func TestPrintArt_EmptyInput(t *testing.T) {
	got := PrintArt("", mockBanner())

	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}
func TestPrintArt_OnlyNewline(t *testing.T) {
	got := PrintArt("\\n", mockBanner())

	if got != "\n" {
		t.Errorf("expected one newline, got %q", got)
	}
}
func TestPrintArt_MultipleNewlines(t *testing.T) {
	got := PrintArt("\\n\\n\\n", mockBanner())

	if got != "\n\n\n" {
		t.Errorf("expected 3 newlines, got %q", got)
	}
}
func TestPrintArt_SingleCharacter(t *testing.T) {
	got := PrintArt("A", mockBanner())

	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")

	if len(lines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(lines))
	}

	if lines[0] != "A1" {
		t.Errorf("expected A1, got %q", lines[0])
	}
}
func TestPrintArt_MultipleCharacters(t *testing.T) {
	got := PrintArt("AB", mockBanner())

	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")

	if lines[0] != "A1B1" {
		t.Errorf("expected A1B1, got %q", lines[0])
	}
}
func TestPrintArt_MultipleLines(t *testing.T) {
	got := PrintArt("A\\nB", mockBanner())

	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")

	if len(lines) != 16 {
		t.Fatalf("expected 16 lines, got %d", len(lines))
	}
}
func TestPrintArt_EmptyLineBetweenText(t *testing.T) {
	got := PrintArt("A\\n\\nB", mockBanner())

	if !strings.Contains(got, "\n\n") {
		t.Errorf("expected blank line between rendered text")
	}
}
func TestPrintArt_LeadingNewline(t *testing.T) {
	got := PrintArt("\\nA", mockBanner())

	if got[0] != '\n' {
		t.Errorf("expected output to start with newline")
	}
}
func TestPrintArt_TrailingNewline(t *testing.T) {
	got := PrintArt("A\\n", mockBanner())

	if !strings.HasSuffix(got, "\n") {
		t.Errorf("expected trailing newline")
	}
}
func TestPrintArt_UnknownCharacter(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("PrintArt panicked on unknown character")
		}
	}()

	PrintArt("@", mockBanner())
}
