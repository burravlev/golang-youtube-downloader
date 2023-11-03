package main

import (
	"fmt"
	youtube "golang-youtube-downloader"
	"log"
)

func main() {
	client := youtube.Client{}
	video, err := client.GetVideo("auD_fT0KCQg")
	if err != nil {
		log.Fatal(err)
	}
	client.Download(youtube.Request{
		Filepath: "saved.mp4",
		Callback: func(percent int) {
			fmt.Println(percent)
		},
		Format: video.StreamingData.Formats[1],
	})
}
