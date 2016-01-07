package engine

import (
	"fmt"
	"log"
	"model"
	"os"
	"path/filepath"
	"strings"
)

var Medias []model.Media

func Scan(path string, fileExtension model.FileExtension) []model.Media {
	var medias []model.Media

	for _, fe := range fileExtension.Extensions {
		//fmt.Println(fe)
		pattern := path + "*" + fe
		files, _ := filepath.Glob(pattern)
		for _, f := range files {
			medias = append(medias, model.Media{FileType: filepath.Ext(f), Name: strings.TrimSuffix(filepath.Base(f), filepath.Ext(f)), Path: f})
		}
	}
	fmt.Println("Scanning directory ", path)
	return medias
}

func ScanFile(fileExtension model.FileExtension) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {

		//fmt.Println("exploring path ",path)
		if err != nil {
			log.Print(err)
			return nil
		}
		if !info.IsDir() {
			f := filepath.Base(path)
			for _, d := range fileExtension.Extensions {
				if strings.HasSuffix(f, d) {
					Medias = append(Medias, model.Media{FileType: filepath.Ext(f), Name: strings.TrimSuffix(filepath.Base(f), filepath.Ext(f)), Path: f})
				}
			}
		} /*else {
			filepath.Walk(path,ScanFile(fileExtension))
		}   */

		return nil
	}
}
