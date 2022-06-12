package factorymethod

import "fmt"

type formatType string

const (
	audioFormat formatType = "Audio"
	videoFormat formatType = "Video"
)

func GetFormat(format string) (IFormat, error) {
	if format == string(audioFormat) {
		return newAudio(), nil
	}
	if format == string(videoFormat) {
		return newVideo(), nil
	}
	return nil, fmt.Errorf("Wrong format type passed")
}
