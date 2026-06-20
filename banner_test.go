package main

import (
	"os"
	"testing"
)

func TestLoadBanner_ValidFile(t *testing.T) {
	content := "\n" +
		"a1\na2\na3\na4\na5\na6\na7\na8\n" +
		"\n" +
		"b1\nb2\nb3\nb4\nb5\nb6\nb7\nb8\n"

	tmp, err := os.CreateTemp("", "banner*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmp.Close()

	banner, err := LoadBanner(tmp.Name())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(banner) != 2 {
		t.Fatalf("expected 2 characters, got %d", len(banner))
	}

	if banner[' '][0] != "a1" {
		t.Errorf("expected first line of space character to be a1")
	}

	if banner['!'][0] != "b1" {
		t.Errorf("expected first line of ! character to be b1")
	}
}

func TestLoadBanner_FileNotFound(t *testing.T) {
	_, err := LoadBanner("does-not-exist.txt")

	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestLoadBanner_EmptyFile(t *testing.T) {
	tmp, err := os.CreateTemp("", "empty*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	tmp.Close()

	_, err = LoadBanner(tmp.Name())

	if err == nil {
		t.Fatal("expected error for empty file")
	}
}

func TestLoadBanner_IncompleteBlock(t *testing.T) {
	content := "\n" +
		"a1\na2\na3\n"

	tmp, err := os.CreateTemp("", "incomplete*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmp.Close()

	_, err = LoadBanner(tmp.Name())

	if err == nil {
		t.Fatal("expected error for incomplete banner block")
	}
}

func TestLoadBanner_CorrectLineCount(t *testing.T) {
	content := "\n" +
		"1\n2\n3\n4\n5\n6\n7\n8\n"

	tmp, err := os.CreateTemp("", "lines*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmp.Close()

	banner, err := LoadBanner(tmp.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(banner[' ']) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(banner[' ']))
	}
}
