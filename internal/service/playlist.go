package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Playlist struct {
	Name   string      `json:"name"`
	Musics []MusicFile `json:"musics"`
}

type PlaylistData struct {
	Playlists []Playlist
	path      string
}

func playlistFilePath() (string, error) {
	var dataDir string

	switch runtime.GOOS {
	case "windows":
		dataDir = os.Getenv("LOCALAPPDATA")
		if dataDir == "" {
			return "", fmt.Errorf("LOCALAPPDATA is not set")
		}
	case "linux", "darwin", "openbsd":
		dataDir = os.Getenv("XDG_DATA_HOME")
		if dataDir == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			dataDir = filepath.Join(home, ".local", "share")
		}
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	dir := filepath.Join(dataDir, "audiogo")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "playlists.json"), nil
}

func LoadPlaylists() (*PlaylistData, error) {
	path, err := playlistFilePath()
	if err != nil {
		return nil, err
	}

	pd := &PlaylistData{path: path}

	b, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return pd, nil
	}
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return pd, nil
	}

	if err := json.Unmarshal(b, &pd.Playlists); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return pd, nil
}

func (pl *PlaylistData) Save() error {
	b, err := json.MarshalIndent(pl.Playlists, "", "  ")
	if err != nil {
		return err
	}

	// write to temp file then rename, so a crash mid-write
	// can't leave you with a corrupted playlists.json
	tmp := pl.path + ".tmp"
	if err := os.WriteFile(tmp, b, 0644); err != nil {
		return err
	}
	return os.Rename(tmp, pl.path)
}

func (pl *PlaylistData) CreateNewPlaylist(name string) {
	pl.Playlists = append(pl.Playlists, Playlist{Name: name})
}

func (pl *PlaylistData) AddMusic(playlistName string, music *MusicFile) error {
	for i := range pl.Playlists {
		if pl.Playlists[i].Name == playlistName {
			pl.Playlists[i].Musics = append(pl.Playlists[i].Musics, *music)
			return nil
		}
	}
	return fmt.Errorf("playlist %q not found", playlistName)
}
