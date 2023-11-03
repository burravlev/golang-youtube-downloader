package youtube

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const API_KEY = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"

var client = &http.Client{}

type Clients struct {
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
	data := []byte(fmt.Sprintf(body, "auD_fT0KCQg"))
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
	return &responseBody, err
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
	file, err := os.Create(request.Filepath)
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
