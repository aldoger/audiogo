package config

import (
	"os"
)

type NodeMusic struct {
	Music         os.DirEntry
	NextNodeMusic *NodeMusic
}

func NewMusic(music os.DirEntry) *NodeMusic {
	return &NodeMusic{
		Music: music,
	}
}

type MusicQueue struct {
	NodeMusicHead *NodeMusic
	NodeMusicTail *NodeMusic
}

func NewMusicQueue() MusicQueue {
	return MusicQueue{
		NodeMusicHead: nil,
		NodeMusicTail: nil,
	}
}

func (q *MusicQueue) Enqueue(music os.DirEntry) {

	newMusic := NewMusic(music)
	if q.NodeMusicHead == nil {
		q.NodeMusicHead = newMusic
		q.NodeMusicTail = newMusic
		return
	}

	q.NodeMusicTail.NextNodeMusic = newMusic
	q.NodeMusicTail = newMusic
}

func (q *MusicQueue) ListMusicInQueue() []string {
	var list []string
	for curr := q.NodeMusicHead; curr != nil; curr = curr.NextNodeMusic {
		list = append(list, curr.Music.Name())
	}
	return list
}
