//go:build prototype

package parser

func crawl(n *html.Node) {
	sch := Schedule{}
	if n.Type == html.ElementNode && n.Data == "div" {
		// entering div
		for _, a := range n.Attr {
			// searching for card
			if a.Key == "class" &&
				a.Val == "card" {
				// searching for card-header Date and card-body Lessons
				var crawlText func(n *html.Node)
				crawlText = func(n *html.Node) {
					if n.Type == html.TextNode {

					}
					c := n.FirstChild
					for c != nil {
						crawlText(c)
						c = c.NextSibling
					}
				}
				crawlText(n)
				break
			}
		}
	}

	c := n.FirstChild
	for c != nil {
		crawl(c)
		c = c.NextSibling
	}
}
