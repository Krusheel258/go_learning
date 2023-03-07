package Zip

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
)

type Archiver interface {
	Archive(names []string, readers ...io.Reader) (outR io.Reader, err error)
}

type Zip struct{}

func New() (z *Zip) {
	return &Zip{}
}

func (z *Zip) Archive(names []string, readers ...io.Reader) (outR io.Reader, err error) {
	ZR, ZW := io.Pipe()
	fmt.Println("pipe created")
	go func() {
		writer := zip.NewWriter(ZW)
		defer ZW.Close()
		fmt.Println("writer created")
		for i, reader := range readers {
			if reader == nil {
				continue
			}
			if names[i] == "" {
				names[i] = fmt.Sprintf("file_%d", i)
			}
			w, _ := writer.Create(names[i])
			log.Println(io.Copy(w, reader))
		}
		writer.Close()
	}()
	return ZR, nil

}
