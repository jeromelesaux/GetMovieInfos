package main

import (
	"client"
	"engine"
	"flag"
	"fmt"
	"model"
	"path/filepath"
)

func main() {

	directoryScanned := flag.String("srcpath", "./", "source path to scan")
	flag.Parse()
	fmt.Println("directory to scan:", *directoryScanned)
	conf := model.LoadConfiguration("extension-file.json")
	filepath.Walk(*directoryScanned, engine.ScanFile(conf))
	fmt.Printf("Found %d medias.\n", len(engine.Medias))
	for _, m := range engine.Medias {
		client.RetrieveInformations(conf, m)
	}
}
