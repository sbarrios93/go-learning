package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Frances", "French")
		want := "Bonjour, Frances"
		assertCorrectMessage(t, got, want)
	})
	t.Run("[SWITCH] saying hello to people", func(t *testing.T) {
		got := HelloSwitch("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("[SWITCH] say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := HelloSwitch("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("[SWITCH] in Spanish", func(t *testing.T) {
		got := HelloSwitch("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("[SWITCH] in French", func(t *testing.T) {
		got := HelloSwitch("Frances", "French")
		want := "Bonjour, Frances"
		assertCorrectMessage(t, got, want)
	})
}
