package parser

import (
	"io"

	"golang.org/x/net/html"
)

type Link struct{
	Href string
	text string
}

func parse(r io.Reader) ([]Link, error){
	doc, err:= html.Parse(r)
	if err!=nil{
		return nil,err
	}

	nodes:= linkNodes(doc)
	var ret[]Link

	for _, node:= range nodes{
		ret = append(ret, buildLink(node))
	}

	return ret,nil
}

func linkNodes(n *html.Node) []*html.Node{
	if n.Type == html.ElementNode || n.Data=="a"{
		return []*html.Node{n}
	}

	var ret []*html.Node
	for ch:= n.FirstChild; ch!= nil; ch= ch.NextSibling{
		ret= append(ret, linkNodes(ch)...)
	}

	return ret
}

func buildLink(n *html.Node) Link{
	var ret Link
	for _, link:= range n.Attr{
		if link.Key == "href"{
			ret.Href=link.Val
			break
		}
	}
	ret.text= extractText(n)
	return ret
}

func extractText(n *html.Node) string{
	if n.Type==html.TextNode{
		return n.Data
	}

	if n.Type!= html.ElementNode{
		return ""
	}

	var ret string
	for ch:= n.FirstChild; ch!= nil; ch=ch.NextSibling{
		ret+= extractText(ch)+" "
	}

	return ret
}