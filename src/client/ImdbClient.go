package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"model"
	"net/http"
	"regexp"
	"strconv"
)

func RetrieveInformations(config model.FileExtension, media model.Media) model.Media {

	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	fmt.Println(media)
	titleToSearch := reg.ReplaceAllString(media.Name, "%20")
	for i := 1; i < 100; i++ {
		requestUrl := config.ImdbUrl + "s=*" + titleToSearch + "*&page=" + strconv.Itoa(i) + "&r=json&plot=full"
		fmt.Println("Get response from ", requestUrl)
		response, err := http.Get(requestUrl)
		if err != nil {
			fmt.Println("Error while getting ", requestUrl, err.Error())
		}
		defer response.Body.Close()
		body, errRead := ioutil.ReadAll(response.Body)
		if errRead != nil {
			fmt.Println("Error while attempting response ", errRead.Error())
		}
		movieSearch := model.ImdbResponse{}
		errDecode := json.Unmarshal([]byte(body), &movieSearch)
		if errDecode != nil {
			fmt.Println("Error while parsing response ", errDecode.Error())
		}
		if len(movieSearch.Search) == 0 {
			break
		}
		media.Information = append(media.Information, movieSearch.Search...)
	}

	fmt.Println(media)

	return media
}
