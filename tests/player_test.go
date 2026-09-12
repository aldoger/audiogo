package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/aldoger/audiogo/internal/service"
)

func TestAudioPlayer(t *testing.T) {
	player := service.NewAudioPlayer()

	fmt.Println("BEFORE PLAY")

	start := time.Now()

	duration, err := player.Play("test-valid.mp3")

	elapsed := time.Since(start)

	fmt.Println("AFTER PLAY")
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Play() elapsed: %v\n", elapsed)

	if err != nil {
		t.Fatalf("Play() returned an error: %v", err)
	}

	if elapsed > time.Second {
		t.Fatalf("Play() appears to be blocking: took %v", elapsed)
	}
}

func TestAudioDuration(t *testing.T) {
	player := service.NewAudioPlayer()

	start := time.Now()

	duration, err := player.Play("test-valid.mp3")

	elapsed := time.Since(start)

	var dur any = duration

	if _, ok := dur.(time.Duration); !ok {
		t.Fatalf("Not type duration")
	}

	if err != nil {
		t.Fatalf("Play() returned an error: %v", err)
	}

	if elapsed > time.Second {
		t.Fatalf("Play() appears to be blocking: took %v", elapsed)
	}
}
