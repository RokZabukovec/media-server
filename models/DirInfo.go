package models

type DirInfo struct {
	Name      string
	Size      int64
	Thumbnail string
	Playlist  string
}

func NewDirInfo(name string, size int64, thumbnail string, playlist string) *DirInfo {
	return &DirInfo{Name: name, Size: size, Thumbnail: thumbnail, Playlist: playlist}
}
