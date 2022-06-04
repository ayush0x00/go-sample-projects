package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl= template.Must(template.ParseFiles("template.html"))
}

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
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct{
	Text string `json:"text"`
	Chapter string `json:"arc"`
}

type handler struct{
	s Story
}

func NewHandler(s Story) http.Handler{
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	path:= r.URL.Path
	
	if path=="" || path=="/"{
		path="/intro"
	}
	path=path[1:]

	if chapter,ok:= h.s[path]; ok{
		err:= tpl.Execute(w, chapter)
		if err!=nil{
		log.Fatalf("Error: %s",err)
		http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w,"Could not find resource", http.StatusNotFound)
}