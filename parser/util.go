package parser

import (
	"fmt"
	"regexp"
	"strings"
)

func containsWord(s, substr string) (match bool) {
	var err error
	for _, field := range strings.Fields(s) {
		match, err = regexp.MatchString(
			fmt.Sprintf("^%s$", substr),
			field,
		)
		if err != nil {
			println(err)
		}
		if match {
			return
		}
	}
	return
}
