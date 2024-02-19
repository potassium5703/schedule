package parser

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const htmlStr = `<div><span id="test">Hello</span><span id="test">World</span></div>`

func TestSearchElem(t *testing.T) {
	doc := Must(
		html.Parse(strings.NewReader(htmlStr)),
	)

	ch := searchElem(doc, "span")
	count := 0
	for range ch {
		count++
	}

	if count != 2 {
		t.Errorf("Expected: 2 span elements, got: %d", count)
	}
}

func TestSearchAttr(t *testing.T) {
	doc := Must(
		html.Parse(strings.NewReader(htmlStr)),
	)

	ch := searchAttr(doc, "id", "test")
	count := 0
	for range ch {
		count++
	}

	if count != 2 {
		t.Errorf("Expected: 2 span elements with id 'test', got: %d", count)
	}
}

func TestSearchElemAttr(t *testing.T) {
	doc := Must(
		html.Parse(strings.NewReader(htmlStr)),
	)

	ch := searchElemAttr(doc, "span", "id", "test")
	count := 0
	for range ch {
		count++
	}

	if count != 2 {
		t.Errorf("Expected: 2 span elements with id 'test', got: %d", count)
	}
}

func TestCrawlTest(t *testing.T) {
	fmt.Println(
		crawlText(Must(
			html.Parse(strings.NewReader(htmlStr)),
		)),
	)
}
