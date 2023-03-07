package Downloader

import (
	"io"
)

type Downloader interface {
	UrlDownload(url string) (r io.Reader, err error)
}

