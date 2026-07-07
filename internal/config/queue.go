package config

type NodeMusic struct {
	Music         string
	NextNodeMusic *NodeMusic
}

func NewMusic(music string) *NodeMusic {
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

func (q *MusicQueue) Enqueue(music string) {

	newMusic := NewMusic(music)
	if q.NodeMusicHead == nil {
		q.NodeMusicHead = newMusic
		q.NodeMusicTail = newMusic
		return
	}

	q.NodeMusicTail.NextNodeMusic = newMusic
	q.NodeMusicTail = newMusic
}

func (q *MusicQueue) Dequeue() string {

	if q.NodeMusicHead == nil && q.NodeMusicTail == nil {
		return ""
	}

	music := q.NodeMusicHead.Music

	currNode := q.NodeMusicHead

	if currNode == q.NodeMusicTail {
		currNode = nil
		q.NodeMusicTail = nil
		return music
	}

	q.NodeMusicHead = q.NodeMusicHead.NextNodeMusic

	currNode = nil

	return music
}

func (q *MusicQueue) ListMusicInQueue() []string {
	var list []string
	for curr := q.NodeMusicHead; curr != nil; curr = curr.NextNodeMusic {
		list = append(list, curr.Music)
	}
	return list
}
