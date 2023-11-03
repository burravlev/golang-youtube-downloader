package main

import (
	"fmt"
	youtube "golang-youtube-downloader"
	"log"
	"os"
)

type test struct {
}

func (*test) Write(b []byte) (int, error) {
	fmt.Println(len(b))
	return len(b), nil
}

type callback struct {
}

func (callback) OnDownloading(i int) {
	fmt.Println(i)
}
func (callback) OnFinished(file *os.File) {
	fmt.Println(file)
}
func (callback) OnError(err error) {
	fmt.Println(err)
}

func main() {
	client := youtube.Client{}
	video, err := client.GetVideo("")
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
