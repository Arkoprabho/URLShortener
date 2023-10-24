package models

import (
	"testing"
)

func TestInvalidity(t *testing.T) {
	url := "a.com/b"
	got := IsValidUrl(url)
	want := false

	if got != want {
		t.Error("Got valid URL when the URL is actually invalid")
	}
}

func TestUrlValidity(t *testing.T) {
	url := "https://a.com/b"
	got := IsValidUrl(url)
	want := true

	if got != want {
		t.Error("Got invalid URL when the URL is actually valid")
	}
}

func TestListTables(t *testing.T) {
	ListTables(cfg)
}
