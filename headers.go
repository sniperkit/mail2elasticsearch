package main

import (
	"regexp"
	"strings"
)

var addrSplitRegex = regexp.MustCompile(`\s*,\s*`)

func splitAddrs(vals []string) []string {
	result := make([]string, 0)
	for _, val := range vals {
		addrs := addrSplitRegex.Split(val, -1)
		result = append(result, addrs...)
	}
	return result
}

var addrRegex = regexp.MustCompile(`[\p{L}\d.!#$%&*+\/=?^_{|}~-]+@[\p{L}\d-.]+`)

func extractOnlyAddrs(vals []string) []string {
	result := make([]string, 0)
	for _, val := range vals {
		result = append(result, addrRegex.FindAllString(val, -1)...)
	}
	return result
}

var whitespaceRegex = regexp.MustCompile(`\s+`)
var commentRegex = regexp.MustCompile(`\([^\)]*\)`)

// RFC 2822 allows whitespace and comments, ElasticSearch/joda-time does not
func stripSpaceAndComments(vals []string) []string {
	result := make([]string, 0)
	for _, val := range vals {
		val = commentRegex.ReplaceAllString(val, "")
		val = whitespaceRegex.ReplaceAllString(val, " ")
		val = strings.TrimSpace(val)
		result = append(result, val)
	}
	return result
}
