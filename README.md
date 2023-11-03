# golang-youtube-downloader
Simple golang parser for retrieving video metadata.

**WARNING**: Youtube API does not support a video download. In fact, it is prohibited - [Terms of Service](https://developers.google.com/youtube/terms/api-services-terms-of-service).

## Usage
### Request
```go
client := youtube.Client{}
video, err := client.GetVideo("abc12345678")
```
### Download
```go
file, err := client.Download(youtube.Request{
	Filepath: "audio",
	Format:   &video.Formats[0],
	// on downloading callback function
	Callback: func(percent int) {
		fmt.Println(percent)
        }   
})
```