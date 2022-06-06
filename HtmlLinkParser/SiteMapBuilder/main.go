package sitemapbuilder

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	parser "github.com/ayush0x00/Parser"
)

func CreateSiteMap(urlPath string) [] string{
	resp,err:= http.Get(urlPath)
	
	if err!=nil{
		log.Fatalf("%s",err)
	}
	defer resp.Body.Close()
	baseURL:= fmt.Sprintf(resp.Request.URL.Scheme+"://"+resp.Request.URL.Host)
	
	return customFilter(parseHrefs(resp.Body,baseURL), withPrefix(baseURL))
}

func parseHrefs(r io.Reader, baseURL string) [] string{
	var ret []string
	links,_:= parser.Parse(r)
	for _,link:= range links{
		switch{
		case strings.HasPrefix(link.Href,"/"):
			ret = append(ret, baseURL+link.Href)
		case strings.HasPrefix(link.Href,"http"):
			ret=append(ret, link.Href)
		}
	}
	return ret
}

func customFilter(links []string, sortFn func(string) bool) [] string{
	var ret []string
	for _, link:= range links{
		if sortFn(link){
			ret=append(ret, link)
		}
	}
	return ret
}

func withPrefix(prfx string) func(string) bool{
	return func(link string) bool{
		return strings.HasPrefix(link,prfx)
	}
}
