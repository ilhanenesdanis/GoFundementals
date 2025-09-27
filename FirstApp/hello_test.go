package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("İlhan")
	want := "Hello İlhan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
