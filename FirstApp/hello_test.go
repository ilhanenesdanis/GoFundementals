package firstapp

import "testing"

func TestHello(t *testing.T) {
	// got := Hello("İlhan")
	// want := "Hello İlhan"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("İlhan", "")
		want := "Hello İlhan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in turkish", func(t *testing.T) {
		got := Hello("İlhan", "Turkish")
		want := "Merhaba İlhan"
		assertCorrectMessage(t, got, want)
	})
}
