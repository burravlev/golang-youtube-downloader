package youtube

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const API_KEY = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"

var client = &http.Client{}

type Clients struct {
}

func (Client) GetVideo(url string) (*VideoInfo, error) {
	return parse(url)
}

func parse(videoId string) (*VideoInfo, error) {
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
	return request(url, "POST", fmt.Sprintf(body, "bZYNGlum3NE"))
}

func request(url, method, body string) (*VideoInfo, error) {
	data := []byte(body)
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
	total uint64
}

func (wc *writeCounter) Write(b []byte) (int, error) {

}

func (Client) Download(format *Format, filepath string, callback YtCallback) {
	file, err := os.Create(filepath)
	if err != nil {
		callback.OnError(err)
		return
	}
	defer file.Close()
	callback.OnFinished(file)
	response, err := http.Get(format.URL)
	if err != nil {
		callback.OnError(err)
		return
	}
	defer response.Body.Close()
}
