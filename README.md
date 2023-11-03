# golang-youtube-downloader
Simple golang parser for retrieving video metadata.

**WARNING**: Youtube API does not support a video download. In fact, it is prohibited - [https://developers.google.com/youtube/terms/api-services-terms-of-service](Terms of Service - II).

## Usage
### Request
```go
client := youtube.Client{}
video, err := client.GetVideo("abc12345678")
```