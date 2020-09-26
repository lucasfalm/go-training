package bubblesort

import "testing"

func TestBubbleSort(t *testing.T) {
	dumbSlice := []int{2, 6, 3, 9, 1, 0, 4}
	BubbleSort(dumbSlice)

	if dumbSlice[0] != 0 {
		t.Errorf("got %d insted 9\n", dumbSlice[0])
	}
}
