package slices

import "testing"

func TestSum(t *testing.T) {

	t.Run("Fixed size", func(t *testing.T) {

		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertMessage(t, got, want)

	})

	t.Run("Dynamic Array or a Slice",func(t *testing.T){

		numbers:= []int{1,2,3}

		got:= Sum((numbers))
		want:= 6

		assertMessage(t,got,want)
	})

}

func assertMessage(t testing.TB, got, want int) {

	if got != want {
		t.Errorf("Got '%d' , Expected '%d'", got, want)
	}

}
