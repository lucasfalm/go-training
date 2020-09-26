package bubblesort

import "testing"

func TestBubbleSort(t *testing.T) {
	dumbSlice := []int{2, 6, 3, 9, 1, 0, 4}
	BubbleSort(dumbSlice)

	if dumbSlice[0] != 0 {
		t.Errorf("got %d insted 0\n", dumbSlice[0])
	}

	if dumbSlice[len(dumbSlice)-1] != 9 {
		t.Errorf("got %d insted 9\n", dumbSlice[len(dumbSlice)-1])
	}
}

func TestMain(t *testing.T) {

}
