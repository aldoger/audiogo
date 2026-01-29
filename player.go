package main

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

type AudioPlayer struct{}

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

}
