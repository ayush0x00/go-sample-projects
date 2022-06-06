package main

import (
	"flag"
	"fmt"

	sitemapbuilder "github.com/ayush0x00/SiteMapBuilder"
)

func main(){
	urlPath:= flag.String("url","https://gophercises.com","url of the site for building sitemap")

	siteMapLinks:= sitemapbuilder.CreateSiteMap(*urlPath, sitemapbuilder.WithPrefix(*urlPath))
	for _, links:= range siteMapLinks{
		fmt.Println(links)
	}
}