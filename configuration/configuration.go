package configuration

const (
	AppName       = "hub"
	Port          = 8888
	MaxFileSize   = 1 << 30         // 1 gigabyte (GB) in bytes
	SegmentSize   = 1024 * 1024 * 5 // 5MB segment size
	BufferSize    = 1024 * 1024     // 1MB buffer size
	SegmentName   = "segment"       // the name of the segment prefix
	PlaylistName  = "manifest.m3u8"
	ThumbnailName = "thumbnail.jpg" // the name of the playlist file
)
