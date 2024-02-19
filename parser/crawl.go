package parser

import (
	"strings"

	"golang.org/x/net/html"
)

type crawlFunc func(*html.Node)

func searchElem(n *html.Node, data string) chan *html.Node {
	ch := make(chan *html.Node)
	var crawl crawlFunc
	crawl = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == data {
			ch <- n
		}
		c := n.FirstChild
		for c != nil {
			crawl(c)
			c = c.NextSibling
		}
	}
	go func() {
		defer close(ch)
		crawl(n)
	}()

	return ch
}

func searchAttr(n *html.Node, key, contains string) chan *html.Node {
	ch := make(chan *html.Node)
	var crawl crawlFunc
	crawl = func(n *html.Node) {
		for _, a := range n.Attr {
			if a.Key == key &&
				containsWord(a.Val, contains) {
				ch <- n
			}
		}
		c := n.FirstChild
		for c != nil {
			crawl(c)
			c = c.NextSibling
		}
	}
	go func() {
		defer close(ch)
		crawl(n)
	}()
	return ch
}

func searchElemAttr(n *html.Node, elem, key, value string) chan *html.Node {
	ch := make(chan *html.Node)
	go func() {
		defer close(ch)
		for e := range searchElem(n, elem) {
			// If document is too large there are
			// would be a hundreds of goroutines :((
			for attr := range searchAttr(e, key, value) {
				ch <- attr
			}
		}
	}()
	return ch
}

func crawlText(n *html.Node) string {
	var s = new(strings.Builder)
	var crawl crawlFunc
	crawl = func(n *html.Node) {
		if n.Type == html.TextNode {
			Must(
				s.WriteString(func() string {
					trimmed := strings.TrimSpace(n.Data) 
					if n.Data == "" {
						return trimmed
					} 
					return trimmed+"\n"
				}()),
			)
		}
		c := n.FirstChild
		for c != nil {
			crawl(c)
			c = c.NextSibling
		}
	}
	crawl(n)
	return s.String()
}
