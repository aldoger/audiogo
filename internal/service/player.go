package service

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

type AudioPlayer struct{}

func NewAudioPlayer() AudioPlayer {
	return AudioPlayer{}
}

func (ap *AudioPlayer) Play(musicFile string) {
	f, err := os.Open(musicFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
