package File

import (
	"ass_3/Downloader"
	"fmt"
	"io"
	"log"
	"os"
)

type filedownload struct {
}

func NewFileDownloader() Downloader.Downloader {
	return &filedownload{}
}

func (fd *filedownload) UrlDownload(url string) (r io.Reader, err error) {
	fptr, err := os.Open(url)
	if err != nil {
		return nil,err
	}
	r, w := io.Pipe()
	go func() {
		defer fptr.Close()
		defer w.Close()
		size, err := io.Copy(w, fptr)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Bytes copied from file to pipe writer : %v", size)
		}
	}()
	return r, nil

}
