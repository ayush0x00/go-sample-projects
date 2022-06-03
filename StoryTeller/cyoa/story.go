package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

func JsonDecoder(r io.Reader) (Story, error){
	d:= json.NewDecoder(r)
	var story Story
	if errDecode:= d.Decode(&story); errDecode!= nil{
		return nil, errDecode
	}
	return story,nil
}

type Chapter struct{
	Title string `json:"title"`
	Paragraphs string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct{
	Text string `json:"text"`
	Chapter string `json:"arc"`
}