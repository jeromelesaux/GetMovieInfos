package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileExtension struct {
	Extensions []string
	ImdbUrl    string
}

func LoadConfiguration(configurationFile string) FileExtension {
	configuration := FileExtension{}
	file, errOpen := os.Open(configurationFile)
	if errOpen != nil {
		fmt.Println("Error while opening file ", configurationFile, errOpen.Error())
	}
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("File extensions supported : ", configuration.Extensions)
	fmt.Println("Imdb API url ", configuration.ImdbUrl)

	return configuration
}
