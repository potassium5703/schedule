package main

import (
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"

	"golang.org/x/net/html"

	"schedule/parser"
)

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

//go:embed config/params.csv
var fs embed.FS

const paramsPath = "config/params.csv"

const tmplString = `{{ range . }}
# {{ .Date }}
{{ range .Lessons }} 
**{{ .What }}**
> {{ .Who }}
> {{ .Where }}
{{ end }}
{{ end }}`

func main() {
	/*
		dateFormat := "2006-01-02"
		now := time.Now().Format(dateFormat)
	*/

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	params := url.Values{}

	{
		const (
			key = iota
			value
		)

		for _, param := range Must(
			csv.NewReader(
				Must(fs.Open(paramsPath)),
			).ReadAll(),
		) {
			params.Add(
				strings.TrimSpace(param[key]), 
				strings.TrimSpace(param[value]),
			)
		}
	}

	/*
		params.Add("scheduler-date", now)
		params.Add("week-up", "week-up") // what coming week come up with?
		params.Add("date-search", now)
	*/

	// fmt.Println(params)

	body := strings.NewReader(params.Encode())

	req := Must(
		http.NewRequest("POST", "https://schedule.ruc.su", body),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp := Must(
		http.DefaultClient.Do(req),
	)
	defer resp.Body.Close()

	/*
		// debug
		Must(
			io.Copy(os.Stdout, resp.Body),
		)
	*/
	/*
		doc := Must(
			html.Parse(resp.Body),
		)

		_ = parser.Parse(doc)
		_ = Must(
			template.New("Schedule").Parse(tmplString),
		)
	*/

	doc := Must(
		html.Parse(io.TeeReader(resp.Body, os.Stdout)),
	)

	scheds := parser.Parse(doc)
	fmt.Println(scheds)
	tmpl := Must(
		template.New("Schedule").Parse(tmplString),
	)
	check(tmpl.Execute(os.Stdout, scheds))
}
