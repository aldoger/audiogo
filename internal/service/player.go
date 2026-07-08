package service

import (
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

type AudioPlayer struct {
	ctrl        *beep.Ctrl
	mixer       *beep.Mixer
	done        chan struct{}
	initialized bool
}

func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{
		mixer: &beep.Mixer{},
	}
}

func (ap *AudioPlayer) Play(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		f.Close()
		return err
	}

	if !ap.initialized {
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

		speaker.Play(ap.mixer)

		ap.initialized = true
	}

	ap.done = make(chan struct{})

	ap.ctrl = &beep.Ctrl{
		Streamer: streamer,
	}

	speaker.Lock()

	ap.mixer.Clear()

	ap.mixer.Add(
		beep.Seq(
			ap.ctrl,
			beep.Callback(func() {
				close(ap.done)
			}),
		),
	)

	speaker.Unlock()

	return nil
}

func (ap *AudioPlayer) Done() <-chan struct{} {
	return ap.done
}

func (ap *AudioPlayer) Pause() {
	if ap.ctrl == nil {
		return
	}

	speaker.Lock()
	ap.ctrl.Paused = true
	speaker.Unlock()
}

func (ap *AudioPlayer) Resume() {
	if ap.ctrl == nil {
		return
	}

	speaker.Lock()
	ap.ctrl.Paused = false
	speaker.Unlock()
}

func (ap *AudioPlayer) Stop() {
	speaker.Lock()
	ap.mixer.Clear()
	speaker.Unlock()

	ap.ctrl = nil
}
