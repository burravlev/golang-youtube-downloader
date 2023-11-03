package youtube

import (
	"fmt"
	"strings"
)

type VideoInfo struct {
	VideoDetails  VideoDetails  `json:"videoDetails"`
	StreamingData StreamingData `json:"streamingData"`
}

type StreamingData struct {
	ExpiresInSeconds string   `json:"expiresInSeconds"`
	Formats          []Format `json:"formats"`
	AdaptiveFormats  []Format `json:"adaptiveFormats"`
}

type Format struct {
	Itag         int    `json:"itag"`
	URL          string `json:"url"`
	MimeType     string `json:"mimeType"`
	Codec        string `json:"-"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	QualityLabel string `json:"qualityLabel"`
	AudioQuality string `json:"audioQuality"`
}

type VideoDetails struct {
	VideoId       string    `json:"videoId"`
	Title         string    `json:"title"`
	LengthSeconds string    `json:"lengthSeconds"`
	Thumbnail     Thumbnail `json:"thumbnail"`
}

type Thumbnail struct {
	Thumbnails []Thumb `json:"thumbnails"`
}

type Thumb struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (v VideoInfo) BestAudioFormat() (*Format, error) {
	formats := v.StreamingData.Formats
	if formats == nil {
		return nil, fmt.Errorf("cannot get audio formats")
	}
	temp := make([]Format, len(formats))
	for _, format := range formats {
		if strings.HasPrefix(format.MimeType, "audio") {
			temp = append(temp, format)
		}
	}
	if v.StreamingData.AdaptiveFormats != nil {
		for _, format := range v.StreamingData.AdaptiveFormats {
			if strings.HasPrefix(format.MimeType, "audio") {
				temp = append(temp, format)
			}
		}
	}
	for _, format := range temp {
		if format.AudioQuality == "AUDIO_QUALITY_HIGH" {
			return &format, nil
		}
	}
	for _, format := range temp {
		if format.AudioQuality == "AUDIO_QUALITY_MEDIUM" {
			return &format, nil
		}
	}
	for _, format := range temp {
		if format.AudioQuality == "AUDIO_QUALITY_LOW" {
			return &format, nil
		}
	}
	return nil, fmt.Errorf("cannot get audio format")
}
