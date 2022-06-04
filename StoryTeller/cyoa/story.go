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
type HandlerOptions func(h *handler)

func WithTemplateOption(t *template.Template) HandlerOptions{
	return func(h *handler) {
		h.t=t
	}
} 

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
	t *template.Template
}

func NewHandler(s Story, opts ...HandlerOptions) http.Handler{
	h:= handler{s,tpl} //default template

	for _, opt:= range opts{
		opt(&h)
	}
	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	path:= r.URL.Path
	
	if path=="" || path=="/"{
		path="/intro"
	}
	path=path[1:]

	if chapter,ok:= h.s[path]; ok{
		err:= h.t.Execute(w, chapter)
		if err!=nil{
		log.Fatalf("Error: %s",err)
		http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w,"Could not find resource", http.StatusNotFound)
}