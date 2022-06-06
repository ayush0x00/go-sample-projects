package sitemapbuilder

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	parser "github.com/ayush0x00/Parser"
)

type FilterOpts func(links *FilteredLinks)
type FilteredLinks struct{
	links []string
}

func CreateSiteMap(urlPath string, opts ...FilterOpts) [] string{
	resp,err:= http.Get(urlPath)
	
	if err!=nil{
		log.Fatalf("%s",err)
	}
	defer resp.Body.Close()
	baseURL:= fmt.Sprintf(resp.Request.URL.Scheme+"://"+resp.Request.URL.Host)
	
	allLinks:= parseHrefs(resp.Body, baseURL)

	for _, filters:= range opts{
		filters(&allLinks)
	}
	return allLinks.links
}

func parseHrefs(r io.Reader, baseURL string) FilteredLinks{
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
	return FilteredLinks{links: ret}
}

func WithPrefix(prfx string) FilterOpts{
	return func(l *FilteredLinks){
		var validLinks []string
		for _,link:= range l.links{
			if strings.HasPrefix(link,prfx){
				validLinks=append(validLinks, link)
			}
		}
		l.links=validLinks
	}
}
