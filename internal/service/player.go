package service

import (
	"errors"
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/flac"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/wav"
)

type AudioFormat string

const (
	MP3  AudioFormat = "mp3"
	WAV  AudioFormat = "wav"
	FLAC AudioFormat = "flac"
	MIDI AudioFormat = "midi"
)

type MusicFile struct {
	Title string
	Path  string
}

type AudioPlayer struct {
	ctrl       *beep.Ctrl
	mixer      *beep.Mixer
	streamer   beep.StreamSeekCloser
	sampleRate beep.SampleRate

	prober *AudioProbe

	done        chan struct{}
	isPaused    bool
	initialized bool
}

var formats = map[string]AudioFormat{
	"mp3":  MP3,
	"wav":  WAV,
	"flac": FLAC,
	"midi": MIDI,
}

func ParseAudioFormat(format string) (AudioFormat, bool) {
	audioFormat, ok := formats[format]
	return audioFormat, ok
}

func InitAudioPlayer() *AudioPlayer {
	prober := InitAudioProbe()

	return &AudioPlayer{
		prober:      prober,
		mixer:       &beep.Mixer{},
		isPaused:    false,
		initialized: false,
	}
}

func (ap *AudioPlayer) decode(format string, file *os.File) (beep.StreamSeekCloser, beep.Format, error) {
	audioFormat, ok := ParseAudioFormat(format)

	if !ok {
		return nil, beep.Format{}, errors.New("format unsupported")
	}

	switch audioFormat {
	case MP3:
		return mp3.Decode(file)
	case WAV:
		return wav.Decode(file)
	case FLAC:
		return flac.Decode(file)
	default:
		return nil, beep.Format{}, errors.New("format unsupported")
	}
}

func (ap *AudioPlayer) Play(file string) (time.Duration, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}

	audioFormat, err := ap.prober.ProbeAudio(file)
	if err != nil {
		f.Close()
		return 0, err
	}

	streamer, beepFormat, err := ap.decode(audioFormat, f)
	if err != nil {
		f.Close()
		return 0, err
	}

	ap.streamer = streamer
	ap.sampleRate = beep.SampleRate(beepFormat.SampleRate)

	samples := streamer.Len()
	duration := beepFormat.SampleRate.D(samples)

	if !ap.initialized {
		speaker.Init(beepFormat.SampleRate, beepFormat.SampleRate.N(time.Second/10))

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

	return duration, nil
}

func (ap *AudioPlayer) CurrentTime() time.Duration {
	if ap.streamer == nil {
		return 0
	}

	speaker.Lock()
	pos := ap.streamer.Position()
	sr := ap.sampleRate
	speaker.Unlock()

	return sr.D(pos)
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
	return ap.initialized
}

func (ap *AudioPlayer) Stop() {
	speaker.Lock()
	ap.mixer.Clear()
	speaker.Unlock()

	ap.ctrl = nil
}
