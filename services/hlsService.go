package services

import (
	"github.com/charmbracelet/log"
	"io"
	"mime/multipart"
)

const (
	SegmentSize = 1024 * 1024 // 5MB segment size
	BufferSize  = 1024        // 1MB buffer size
	SegmentName = "segment"   // the name of the segment prefix
)

func CreateSegmentedFiles(folderName string, file multipart.File) (int64, error) {
	segmentsFolderPath, err := CreateHlsFolder(folderName)
	if err != nil {
		log.Error("Failed to create media folder")
	}

	buffer := make([]byte, BufferSize)
	var nbBytes int64
	segmentIndex := 0
	segmentFile, err := CreateNewSegmentFile(segmentsFolderPath, SegmentName, segmentIndex)

	if err != nil {
		return 0, err
	}

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			return nbBytes, err
		}

		nbBytes += int64(n)

		if nbBytes >= int64(SegmentSize) {
			segmentIndex++
			segmentFile.Close()
			segmentFile, err = CreateNewSegmentFile(segmentsFolderPath, SegmentName, segmentIndex)

			if err != nil {
				return nbBytes, err
			}

			defer segmentFile.Close()

			nbBytes = 0
		}

		_, writeErr := segmentFile.Write(buffer[:n])

		if writeErr != nil {
			log.Error("Failed to write to file")
		}
	}

	return nbBytes, nil
}
