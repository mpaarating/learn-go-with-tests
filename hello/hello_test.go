package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, actual, expected string) {
		t.Helper()
		if actual != expected {
			t.Errorf("actual %q expected %q", actual, expected)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Chris", "")
		expected := "Hello, Chris"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		actual := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in French", func(t *testing.T) {
		actual := Hello("Sammie", "French")
		expected := "Bonjour, Sammie"
		assertCorrectMessage(t, actual, expected)
	})

}
