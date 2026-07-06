package utils

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func DirExist() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dirPath := filepath.Join(homeDir, "Music")

	_, err = os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	return dirPath, nil
}

func ListMusic(path string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	audioExt := map[string]bool{
		".mp3":  true,
		".wav":  true,
		".flac": true,
	}

	var musicFiles []os.DirEntry

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if audioExt[ext] {
			musicFiles = append(musicFiles, file)
		}
	}

	if len(musicFiles) < 1 {
		return nil, errors.New("empty directory, no audio files found")
	}

	return musicFiles, nil
}
