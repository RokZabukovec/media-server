package services

import (
	"fmt"
	"github.com/charmbracelet/log"
	"io"
	"mediaserver/configuration"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func CreateHLSFilesFromAPIRequest(file multipart.File, outputFolderPath, outputFilename string) error {
	if _, err := os.Stat(outputFolderPath); os.IsNotExist(err) {
		if err = os.MkdirAll(outputFolderPath, os.ModePerm); err != nil {
			log.Errorf("failed to create output directory: %w", err)
			return err
		}
	}

	// Create or overwrite the output video file
	outputPath := filepath.Join(outputFolderPath, outputFilename)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Errorf("Failed to create output file: %w", err)
		return err
	}

	defer outputFile.Close()
	defer func() {
		if err = os.Remove(outputPath); err != nil {
			log.Errorf("Failed to remove temporary file: %w", err)
		}
	}()

	// Copy the content of the uploaded file to the output file
	if _, err = io.Copy(outputFile, file); err != nil {
		log.Errorf("Failed to copy uploaded file to output file: %w", err)
		return err
	}

	// Create the HLS playlist from the video
	hlsOutputPath := filepath.Join(outputFolderPath, configuration.PlaylistName)
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i", outputPath,
		"-profile:v", "baseline",
		"-level", "3.0",
		"-start_number", "0",
		"-hls_time", "10",
		"-hls_list_size", "0",
		"-f", "hls",
		"-preset", "ultrafast",
		hlsOutputPath,
	)

	output, err := ffmpegCmd.CombinedOutput()
	if err != nil {
		log.Errorf("Failed to create HLS: %v\nOutput: %s", err, string(output))
		return err
	}

	thumbnailOutputPath := filepath.Join(outputFolderPath, configuration.ThumbnailName)

	// TODO Add check if the video is less than 5s long.
	// Take the frame in the middle
	if _, err := os.Stat(thumbnailOutputPath); os.IsNotExist(err) {
		thumbnailCmd := exec.Command(
			"ffmpeg",
			"-i", outputPath,
			"-ss", "00:00:05",
			"-vframes", "1",
			thumbnailOutputPath,
		)

		if output, err = thumbnailCmd.CombinedOutput(); err != nil {
			log.Errorf("Failed to extract thumbnail: %v\nOutput: %s", err, string(output))
			return err
		}
	}

	return nil
}

// CreateSegmentedFiles An attempt to segment the video manually
// without the use of external dependencies like FFMPEG.
// It proved to have too many unknowns to be viable
// for now. Don't reinvent the wheel.
func CreateSegmentedFiles(folderName string, file multipart.File) (int64, error) {
	segmentsFolderPath, err := CreateHlsFolder(folderName)
	if err != nil {
		log.Error("Failed to create media folder")
	}

	buffer := make([]byte, configuration.BufferSize)
	segmentIndex := 0
	segmentFile, err := CreateNewSegmentFile(segmentsFolderPath, configuration.SegmentName, segmentIndex)
	defer func(segmentFile *os.File) {
		err := segmentFile.Close()
		if err != nil {
			log.Errorf("Could not close the file %s", segmentFile.Name())
		}
	}(segmentFile)

	if err != nil {
		return 0, err
	}
	var nbBytes int64
	var totalBytes int64
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			return nbBytes, err
		}

		nbBytes += int64(n)
		totalBytes += int64(n)

		if nbBytes >= int64(configuration.SegmentSize) {
			segmentIndex++
			segmentFile.Close()
			segmentFile, err = CreateNewSegmentFile(segmentsFolderPath, configuration.SegmentName, segmentIndex)

			if err != nil {
				return nbBytes, err
			}

			nbBytes = 0
		}

		_, writeErr := segmentFile.Write(buffer[:n])

		if writeErr != nil {
			log.Error("Failed to write to file")
		}
	}

	_ = GenerateHLSManifest(segmentsFolderPath, 10.0)

	return totalBytes, nil
}

func GenerateHLSManifest(filepath string, segmentDuration float64) error {
	file, err := os.Create(path.Join(filepath, "manifest.m3u8"))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("#EXTM3U\n")
	_, err = file.WriteString("#EXT-X-VERSION:3\n")
	_, err = file.WriteString("#EXT-X-TARGETDURATION:15\n")
	_, err = file.WriteString("#EXT-X-MEDIA-SEQUENCE:0\n")

	if err != nil {
		return err
	}

	segmentFiles, err := os.ReadDir(filepath)
	if err != nil {
		return err
	}

	for _, segmentInfo := range segmentFiles {
		if segmentInfo.IsDir() {
			continue
		}
		segmentURI := segmentInfo.Name()
		_, err = file.WriteString(fmt.Sprintf("#EXTINF:%.3f,\n%s\n", segmentDuration, segmentURI))
		if err != nil {
			return err
		}
	}

	_, err = file.WriteString("#EXT-X-ENDLIST\n")

	return nil
}
