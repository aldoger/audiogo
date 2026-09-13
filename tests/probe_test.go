package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aldoger/audiogo/internal/service"
)

func TestAudioProbe_ProbeAudio(t *testing.T) {
	probe := service.InitAudioProbe()

	tests := []struct {
		name           string
		file           string
		expectedFormat string
	}{
		{
			name:           "MP3",
			file:           "test-valid.mp3",
			expectedFormat: "mp3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := filepath.Abs(tt.file)
			if err != nil {
				t.Fatal(err)
			}

			if _, err := os.Stat(path); err != nil {
				t.Skipf("test file does not exist: %s", path)
			}

			format, err := probe.ProbeAudio(path)
			if err != nil {
				t.Fatalf("ProbeAudio() error = %v", err)
			}

			if format != tt.expectedFormat {
				t.Errorf(
					"ProbeAudio() = %q, want %q",
					format,
					tt.expectedFormat,
				)
			}
		})
	}
}
