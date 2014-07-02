package data

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileImporter struct {
	basePath string
	importFn StoreFunc
}

func NewFileImporter(basePath string) *FileImporter {
	return &FileImporter{
		basePath: basePath,
	}
}

func (fi *FileImporter) SetStoreFunc(fn StoreFunc) {
	fi.importFn = fn
}

func (fi FileImporter) ImportData() {
	dirs := []typeDir{VideoDir, SpeakerDir, EventDir, EventSeriesDir}
	done := make(chan string)

	for _, dir := range dirs {
		d := dir

		walkFn := func(path string, info os.FileInfo, err error) error {
			f, err := os.Open(path)
			if err != nil {
				log.Printf("open err: %s\n", err)
				return err
			}
			return imp(f, d, fi.importFn)
		}

		go func() {
			filepath.Walk(typeDirPath(fi.basePath, d), filterFiles(walkFn))
			done <- d.String()
		}()
	}

	// wait for all walks to finish
	c := 0
L:
	for {
		select {
		case s := <-done:
			c++
			log.Println("finished: ", s)
			if c == len(dirs) {
				break L
			}
		}
	}
}

func filterFiles(f func(path string, info os.FileInfo, err error) error) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if info == nil {
			s := fmt.Sprintf("info is nil (%s)", path)
			log.Println(s)
			return fmt.Errorf(s)
		}
		if !info.IsDir() && strings.HasSuffix(path, ".json") {
			return f(path, info, err)
		}
		return nil
	}
}
