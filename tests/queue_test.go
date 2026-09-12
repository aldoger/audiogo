package tests

import (
	"testing"

	"github.com/aldoger/audiogo/internal/service"
)

func TestNewMusicQueue(t *testing.T) {
	q := service.NewMusicQueue()

	if q.NodeMusicHead != nil {
		t.Fatal("expected head to be nil")
	}

	if q.NodeMusicTail != nil {
		t.Fatal("expected tail to be nil")
	}
}

func TestEnqueue(t *testing.T) {
	q := service.NewMusicQueue()

	q.Enqueue("song1.mp3")

	if q.NodeMusicHead == nil {
		t.Fatal("expected head to be non-nil")
	}

	if q.NodeMusicTail == nil {
		t.Fatal("expected tail to be non-nil")
	}

	if q.NodeMusicHead.Music != "song1.mp3" {
		t.Fatalf("expected head music to be song1.mp3, got %s", q.NodeMusicHead.Music)
	}

	if q.NodeMusicTail.Music != "song1.mp3" {
		t.Fatalf("expected tail music to be song1.mp3, got %s", q.NodeMusicTail.Music)
	}
}

func TestEnqueueMultiple(t *testing.T) {
	q := service.NewMusicQueue()

	q.Enqueue("song1.mp3")
	q.Enqueue("song2.mp3")
	q.Enqueue("song3.mp3")

	expected := []string{
		"song1.mp3",
		"song2.mp3",
		"song3.mp3",
	}

	actual := q.ListMusicInQueue()

	if len(actual) != len(expected) {
		t.Fatalf("expected %d songs, got %d", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Fatalf(
				"expected song at index %d to be %s, got %s",
				i,
				expected[i],
				actual[i],
			)
		}
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := service.NewMusicQueue()

	music := q.Dequeue()

	if music != "" {
		t.Fatalf("expected empty string, got %s", music)
	}
}

func TestDequeueSingle(t *testing.T) {
	q := service.NewMusicQueue()

	q.Enqueue("song1.mp3")

	music := q.Dequeue()

	if music != "song1.mp3" {
		t.Fatalf("expected song1.mp3, got %s", music)
	}

	if q.NodeMusicHead != nil {
		t.Fatal("expected head to be nil after dequeue")
	}

	if q.NodeMusicTail != nil {
		t.Fatal("expected tail to be nil after dequeue")
	}
}

func TestDequeueFIFO(t *testing.T) {
	q := service.NewMusicQueue()

	q.Enqueue("song1.mp3")
	q.Enqueue("song2.mp3")
	q.Enqueue("song3.mp3")

	expected := []string{
		"song1.mp3",
		"song2.mp3",
		"song3.mp3",
	}

	for i, expectedMusic := range expected {
		music := q.Dequeue()

		if music != expectedMusic {
			t.Fatalf(
				"dequeue %d: expected %s, got %s",
				i,
				expectedMusic,
				music,
			)
		}
	}
}

func TestDequeueUntilEmpty(t *testing.T) {
	q := service.NewMusicQueue()

	q.Enqueue("song1.mp3")
	q.Enqueue("song2.mp3")

	q.Dequeue()
	q.Dequeue()

	if q.NodeMusicHead != nil {
		t.Fatal("expected head to be nil")
	}

	if q.NodeMusicTail != nil {
		t.Fatal("expected tail to be nil")
	}

	if music := q.Dequeue(); music != "" {
		t.Fatalf("expected empty queue, got %s", music)
	}
}

func TestListMusicInQueueEmpty(t *testing.T) {
	q := service.NewMusicQueue()

	list := q.ListMusicInQueue()

	if len(list) != 0 {
		t.Fatalf("expected empty list, got %d items", len(list))
	}
}

func TestListMusicInQueue(t *testing.T) {
	q := service.NewMusicQueue()

	q.Enqueue("song1.mp3")
	q.Enqueue("song2.mp3")
	q.Enqueue("song3.mp3")

	list := q.ListMusicInQueue()

	expected := []string{
		"song1.mp3",
		"song2.mp3",
		"song3.mp3",
	}

	if len(list) != len(expected) {
		t.Fatalf("expected %d items, got %d", len(expected), len(list))
	}

	for i := range expected {
		if list[i] != expected[i] {
			t.Fatalf(
				"index %d: expected %s, got %s",
				i,
				expected[i],
				list[i],
			)
		}
	}
}
