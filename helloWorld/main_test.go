package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Greeting People", func(t *testing.T) {
		got := HelloWorld("Kannav", "")
		want := "Hello, Kannav"

		assertMessage(t, got, want)

	})

	t.Run("Sending Hello World if no input is given", func(t *testing.T) {
		got := HelloWorld("", "")
		want := "Hello, World"

		assertMessage(t, got, want)

	})
	t.Run("User Speaks Spanish", func(t *testing.T) {
		got := HelloWorld("Kannav", "Spanish")
		want := "Hola, Kannav"

		assertMessage(t, got, want)
	})

	t.Run("User Speaks French", func(t *testing.T) {
		got := HelloWorld("Kannav", "French")
		want := "Bonjour, Kannav"

		assertMessage(t, got, want)
	})

}
func assertMessage(t testing.TB, got, want string) {
	// used to specify that this function is a helper function for the other tests
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
