package service

type Playlist struct {
	Title  string
	Musics []MusicFile
}

func CreatePlaylist(title string) *Playlist {
	return &Playlist{
		Title:  title,
		Musics: nil,
	}
}

func GetPlaylistData() {

}

func (pl *Playlist) AddMusic()
