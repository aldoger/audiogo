package service

import (
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

type MusicFile struct {
	Title string
	Path  string
}

type AudioPlayer struct {
	ctrl        *beep.Ctrl
	mixer       *beep.Mixer
	done        chan struct{}
	isPaused    bool
	initialized bool
}

func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{
		mixer:       &beep.Mixer{},
		isPaused:    false,
		initialized: false,
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
	ap.isPaused = true
	speaker.Unlock()
}

func (ap *AudioPlayer) Resume() {
	if ap.ctrl == nil {
		return
	}

	speaker.Lock()
	ap.ctrl.Paused = false
	ap.isPaused = false
	speaker.Unlock()
}

func (ap *AudioPlayer) IsPaused() bool {
	return ap.isPaused
}

func (ap *AudioPlayer) IsInitialized() bool {
	return ap.IsInitialized()
}

func (ap *AudioPlayer) Stop() {
	speaker.Lock()
	ap.mixer.Clear()
	speaker.Unlock()

	ap.ctrl = nil
}
