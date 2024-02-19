package parser

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"golang.org/x/net/html"
)

type Lesson struct {
	What, Who, Where string
}

func (l Lesson) String() string {
	return fmt.Sprint(
		"What:", l.What, "\n",
		"Who:", l.Who, "\n",
		"Where:", l.Where, "\n",
	)
}

type Schedule struct {
	Date    string
	Lessons []Lesson
}

func (s Schedule) String() string {
	return fmt.Sprint(
		"Date:", s.Date, "\n",
		"Lessons:", s.Lessons, "\n",
	)
}

func isEOF(err error) bool {
	if err != nil {
		switch err {
		case io.EOF:
			return true
		default:
			panic(err)
		}
	}
	return false
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](v T, err error) T {
	check(err)
	return v
}

func Parse(n *html.Node) []Schedule {
	var scheds []Schedule
	// Debug purposes.
	var count int
	for card := range searchElemAttr(n, "div", "class", "card") {
		// fmt.Print(crawlText(card))
		count++
		sch := Schedule{}
		sch.Date = func() string {
			date := <-searchElemAttr(card, "div", "class", "card-header")
			return crawlText(date)

		}()
		for lesson := range searchElemAttr(card, "div", "class", "card-body") {
			s := func() (res []string) {
				for _, s := range strings.Split(crawlText(lesson), "\n") {
					if s == "" {
						continue
					}
					res = append(res, s)
				}
				return
			}()
			var l Lesson
			v := reflect.ValueOf(&l).Elem()
			for i := 0; i < v.NumField(); i++ {
				v.Field(i).SetString(s[i])
			}

			sch.Lessons = append(sch.Lessons, l)
		}
		scheds = append(scheds, sch)
	}
	println("DEBUG:", count)
	return scheds
}
