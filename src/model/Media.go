package model

type Media struct {
	Name        string
	Path        string
	FileType    string
	Information []MovieInformation
}

type ImdbResponse struct {
	Search []MovieInformation
}

type MovieInformation struct {
	Title  string
	Year   string
	imdbId string
	Type   string
	Poster string
}
