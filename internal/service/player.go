package service

import (
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

type AudioPlayer struct {
	streamer beep.StreamSeekCloser
	ctrl     *beep.Ctrl
	done     chan bool
}

func NewAudioPlayer() AudioPlayer {
	return AudioPlayer{}
}

func (ap *AudioPlayer) Play(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}

	ap.streamer = streamer
	ap.done = make(chan bool)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ap.ctrl = &beep.Ctrl{
		Streamer: streamer,
		Paused:   false,
	}

	speaker.Play(
		beep.Seq(
			ap.ctrl,
			beep.Callback(func() {
				ap.done <- true
			}),
		),
	)

	return nil
}

func (ap *AudioPlayer) Done() <-chan bool {
	return ap.done
}

func (ap *AudioPlayer) Pause() {
	speaker.Lock()
	ap.ctrl.Paused = true
	speaker.Unlock()
}

func (ap *AudioPlayer) Resume() {
	speaker.Lock()
	ap.ctrl.Paused = false
	speaker.Unlock()
}

func (ap *AudioPlayer) Stop() {
	speaker.Clear()

	if ap.streamer != nil {
		ap.streamer.Close()
	}
}
