package feedapp

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})

	if len(feed.Items) == 0 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})

	total := feed.GetAll()

	if len(total) != 1 {
		t.Errorf("Item was not added")
	}
}
