package utils

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/themaluxis/accdiscordbot/pkg/data"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func FormatTime(ms int) string {
	duration := time.Duration(ms) * time.Millisecond
	minutes := duration / time.Minute
	duration -= minutes * time.Minute
	seconds := duration / time.Second
	duration -= seconds * time.Second
	milliseconds := duration / time.Millisecond
	return fmt.Sprintf("%02d:%02d:%03d", minutes, seconds, milliseconds)
}

var MessageIDs = make(map[string]string)

func convertToUTF8AndClean(r io.Reader) (io.Reader, error) {
	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, r)
	if err != nil {
		return nil, err
	}

	// Remove null characters
	cleanedBytes := bytes.Map(func(r rune) rune {
		if r == rune(0) {
			return -1
		}
		return r
	}, buffer.Bytes())

	// Detect and convert encoding
	encoding, _, _ := charset.DetermineEncoding(cleanedBytes, "")
	if encoding == nil {
		return nil, fmt.Errorf("unable to determine encoding")
	}

	return transform.NewReader(bytes.NewReader(cleanedBytes), encoding.NewDecoder()), nil
}

func UpdateBestLap(groupedBestLaps map[string]map[int][]data.BestLapInfo, trackName string, bestLap data.BestLapInfo) {
	if _, exists := groupedBestLaps[trackName]; !exists {
		groupedBestLaps[trackName] = make(map[int][]data.BestLapInfo)
	}

	found := false
	for i, lap := range groupedBestLaps[trackName][bestLap.CarModel] {
		if lap.DriverID == bestLap.DriverID {
			found = true
			if bestLap.BestLap < lap.BestLap {
				groupedBestLaps[trackName][bestLap.CarModel][i] = bestLap
			}
			break
		}
	}
	if !found {
		groupedBestLaps[trackName][bestLap.CarModel] = append(groupedBestLaps[trackName][bestLap.CarModel], bestLap)
	}
}
