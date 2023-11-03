package youtube

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const API_KEY = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"

var client = &http.Client{}

type Client struct {
}

func (Client) GetVideo(videoId string) (*VideoInfo, error) {
	url := "https://youtubei.googleapis.com/youtubei/v1/player?key=" + API_KEY
	body := `{
		"videoId" : "%s",
		"context" : {
			"client" : {
				"hl" : "en",
				"gl" : "US",
				"clientName" : "ANDROID_TESTSUITE",
				"clientVersion" : "1.9",
				"androidSdkVersion" : "31"
			}
		}
	}`
	data := []byte(fmt.Sprintf(body, videoId))
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("api error: cannot make request")
	}
	var responseBody VideoInfo
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	for i, format := range responseBody.StreamingData.Formats {
		parseCodecs(&format)
		responseBody.StreamingData.Formats[i] = format
	}
	return &responseBody, err
}

func parseCodecs(format *Format) {
	arr := strings.Split(format.MimeType, ";")
	if len(arr) > 1 {
		format.MimeType = arr[0]
		format.Codec = strings.TrimSpace(arr[1])
	}
}

type videoInfo struct {
}

type writeCounter struct {
	percen   int
	loaded   int64
	total    int64
	callback func(int)
}

func (wc *writeCounter) Write(b []byte) (int, error) {
	wc.loaded += int64(len(b))
	if current := int(wc.loaded * 100 / wc.total); current > wc.percen {
		wc.percen = current
		wc.callback(wc.percen)
	}
	return len(b), nil
}

func (Client) Download(request Request) (*os.File, error) {
	if len(request.Format.MimeType) <= 6 {
		return nil, fmt.Errorf("cannot extract file extension")
	}
	ext := request.Format.MimeType[6:]
	file, err := os.Create(request.Filepath + "." + ext)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	response, err := http.Get(request.Format.URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contentLength, err := strconv.ParseInt(response.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return nil, err
	}
	if request.Callback != nil {
		wr := &writeCounter{
			percen:   0,
			loaded:   0,
			total:    contentLength,
			callback: request.Callback,
		}
		if _, err := io.Copy(file, io.TeeReader(response.Body, wr)); err != nil {
			return nil, err
		}
		return file, nil
	}
	_, err = io.Copy(file, response.Body)
	return file, err
}
