package service

import (
	"encoding/json"
	"os/exec"
)

type AudioProbe struct{}

type ffprobeOutput struct {
	Format struct {
		FormatName string `json:"format_name"`
	} `json:"format"`
}

func InitAudioProbe() *AudioProbe {
	return &AudioProbe{}
}

func (a *AudioProbe) ProbeAudio(path string) (string, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		path,
	)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	var probe ffprobeOutput

	if err := json.Unmarshal(output, &probe); err != nil {
		return "", err
	}

	return probe.Format.FormatName, nil
}
