package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type SEOReport struct {
	HasTitle       bool   `json:"has_title"`
	Title          string `json:"title"`
	HasDescription bool   `json:"has_description"`
	Description    string `json:"description"`
	HasH1          bool   `json:"has_h1"`
}

type AssetInfo struct {
	URL  string
	Type string
}

func ParseSEO(r io.Reader) *SEOReport {
	seo := &SEOReport{}

	doc, err := html.Parse(r)
	if err != nil {
		return seo
	}

	var titleText, descriptionText string
	var foundTitle, foundDescription, foundH1 bool
	var inChannel, inItem, inHead bool
	var channelTitleFound bool

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "channel":
				inChannel = true
			case "item":
				inItem = true
			case "head":
				inHead = true
			case "title":
				if n.FirstChild != nil && !foundTitle {
					titleVal := strings.TrimSpace(n.FirstChild.Data)
					if inChannel && !inItem && !channelTitleFound {
						titleText = titleVal
						foundTitle = true
						channelTitleFound = true
					}
					if inHead && !foundTitle {
						titleText = titleVal
						foundTitle = true
					}
					if !inChannel && !inHead && !foundTitle {
						titleText = titleVal
						foundTitle = true
					}
				}
			case "meta":
				var name, content string
				for _, attr := range n.Attr {
					if attr.Key == "name" {
						name = strings.ToLower(attr.Val)
					}
					if attr.Key == "content" {
						content = attr.Val
					}
				}
				if name == "description" {
					descriptionText = content
					foundDescription = true
				}
			case "h1":
				foundH1 = true
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}

		switch n.Data {
		case "channel":
			inChannel = false
		case "item":
			inItem = false
		case "head":
			inHead = false
		}
	}

	walk(doc)

	seo.HasTitle = foundTitle
	seo.Title = titleText
	seo.HasDescription = foundDescription
	seo.Description = descriptionText
	seo.HasH1 = foundH1

	return seo
}

func ParseHTMLLinks(r io.Reader) ([]string, error) {
	var links []string

	doc, err := html.Parse(r)
	if err != nil {
		return links, err
	}

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && attr.Val != "" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(doc)

	return links, nil
}

func ParseAssets(r io.Reader) []AssetInfo {
	var assets []AssetInfo

	doc, err := html.Parse(r)
	if err != nil {
		return assets
	}

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "img":
				for _, attr := range n.Attr {
					if attr.Key == "src" && attr.Val != "" {
						assets = append(assets, AssetInfo{URL: attr.Val, Type: "image"})
					}
				}
			case "script":
				for _, attr := range n.Attr {
					if attr.Key == "src" && attr.Val != "" {
						assets = append(assets, AssetInfo{URL: attr.Val, Type: "script"})
					}
				}
			case "link":
				var rel, href string
				for _, attr := range n.Attr {
					if attr.Key == "rel" {
						rel = strings.ToLower(attr.Val)
					}
					if attr.Key == "href" {
						href = attr.Val
					}
				}
				if rel == "stylesheet" && href != "" {
					assets = append(assets, AssetInfo{URL: href, Type: "style"})
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(doc)

	return assets
}
