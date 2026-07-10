package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/aldoger/audiogo/internal/service"
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

func ListMusic(path string) ([]service.MusicFile, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	audioExt := map[string]bool{
		".mp3": true,
	}

	var musicFiles []service.MusicFile

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if audioExt[ext] {
			music := service.MusicFile{Title: file.Name(), Path: filepath.Join(path, file.Name())}
			musicFiles = append(musicFiles, music)
		}
	}

	if len(musicFiles) < 1 {
		return nil, nil
	}

	return musicFiles, nil
}
