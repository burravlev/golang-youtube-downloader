package youtube

type VideoInfo struct {
	VideoDetails  VideoDetails  `json:"videoDetails"`
	StreamingData StreamingData `json:"streamingData"`
}

type StreamingData struct {
	ExpiresInSeconds string   `json:"expiresInSeconds"`
	Formats          []Format `json:"formats"`
}

type Format struct {
	Itag         int    `json:"itag"`
	URL          string `json:"url"`
	MimeType     string `json:"mimeType"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	QualityLabel string `json:"qualityLabel"`
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
