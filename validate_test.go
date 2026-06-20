package main

import "testing"

func TestValidate_EmptyString(t *testing.T) {
	rn, err := Validate("")

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if rn != 0 {
		t.Fatalf("expected rune 0, got %q", rn)
	}
}

func TestValidate_ValidText(t *testing.T) {
	rn, err := Validate("Hello World!")

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if rn != 0 {
		t.Fatalf("expected rune 0, got %q", rn)
	}
}

func TestValidate_SpaceBoundary(t *testing.T) {
	rn, err := Validate(" ")

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if rn != 0 {
		t.Fatalf("expected rune 0, got %q", rn)
	}
}

func TestValidate_TildeBoundary(t *testing.T) {
	rn, err := Validate("~")

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if rn != 0 {
		t.Fatalf("expected rune 0, got %q", rn)
	}
}

func TestValidate_Newline(t *testing.T) {
	rn, err := Validate("\n")

	if err == nil {
		t.Fatal("expected error")
	}

	if rn != '\n' {
		t.Fatalf("expected newline rune, got %q", rn)
	}
}

func TestValidate_Tab(t *testing.T) {
	rn, err := Validate("\t")

	if err == nil {
		t.Fatal("expected error")
	}

	if rn != '\t' {
		t.Fatalf("expected tab rune, got %q", rn)
	}
}

func TestValidate_DELCharacter(t *testing.T) {
	rn, err := Validate(string(rune(127)))

	if err == nil {
		t.Fatal("expected error")
	}

	if rn != rune(127) {
		t.Fatalf("expected rune 127, got %d", rn)
	}
}

func TestValidate_UnicodeCharacter(t *testing.T) {
	rn, err := Validate("é")

	if err == nil {
		t.Fatal("expected error")
	}

	if rn != 'é' {
		t.Fatalf("expected é, got %q", rn)
	}
}

func TestValidate_FirstInvalidCharacterReturned(t *testing.T) {
	input := "Hello\n\tWorld"

	rn, err := Validate(input)

	if err == nil {
		t.Fatal("expected error")
	}

	if rn != '\n' {
		t.Fatalf("expected first invalid rune to be newline, got %q", rn)
	}
}
