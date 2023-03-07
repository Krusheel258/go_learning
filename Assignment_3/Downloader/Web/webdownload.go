package Web

import (
	"ass_3/Downloader"
	"fmt"
	"io"
	"log"
	"net/http"
)

type webdownload struct {
}

func NewWebDownloader() Downloader.Downloader {
	return &webdownload{}
}

func (wd *webdownload) UrlDownload(url string) (r io.Reader, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	r, w := io.Pipe()
	go func() {
		defer resp.Body.Close()
		defer w.Close()
		size, err := io.Copy(w, resp.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Bytes copied from resp.Body to pipe writer : %v", size)
		}
	}()
	return r, nil

}
