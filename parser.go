package ytdl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const API_KEY = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"

var client = &http.Client{}

func parseVideoAndroid(videoId string) {
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
	json, err := parse(url, "POST", body)
}

func parse(url, method, body string) (map[string]json.RawMessage, error) {
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
	var responseBody map[string]json.RawMessage
	err = json.NewDecoder(response.Body).Decode(&body)
	return responseBody, err
}

func parseWeb(url, method, body string) {

}

func parseVideoDetails(videoId string, json map[string]json.RawMessage) {

}

func formats(json map[string]json.RawMessage) ([]Format, error) {
	streamingData, ok := json["streamingData"]
	if !ok {
		return nil, fmt.Errorf("cannot get formats")
	}
}
