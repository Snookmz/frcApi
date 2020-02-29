package core

type Ini struct {
	Api Api `json:"api"`
	Log Log `json:"log"`
	Dir string `json:"dir"`
	DropBox DropBox `json:"dropBox"`
}

type Api struct {
	Port int `json:"port"`
}

type Log struct {
	File string `json:"file"`
}

type DropBox struct {
	AccessToken string `json:"accessToken"`
	Url string `json:"url"`
	Path string `json:"path"`
}