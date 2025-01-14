package iteration

import "testing"

func TestIteration(t *testing.T) {

	t.Run("No iteration is mentioned", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "aaaaa"

		assertMessage(t, got, want)

	})

	t.Run("Iteration is mentioned", func(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"

		assertMessage(t, got, want)

	})

}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q , expected %q", got, want)
	}

}
func BenchmarkIteration(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
