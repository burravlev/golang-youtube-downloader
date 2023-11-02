package youtube

import "os"

type YtCallback interface {
	OnDownloading(int)
	OnFinished(*os.File)
	OnError(error)
}
