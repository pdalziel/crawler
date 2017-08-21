package main

import (
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

// Test CLI commands ("-h", "-url=")
// -h show the available commands
// possible features to add:
//		-f specify output file
//		-v verbose logging

// Test display commands
func TestDisplayMsg(t *testing.T) {
	var commandTests = []struct {
		in  string
		out string
	}{
		{
			in:  "",
			out: "usage: '$ crawl url-to-crawl'",
		},
		{
			in:  "-h",
			out: "Available commands" + "'-h'  :   List available commands",
		},
		{
			in:  "12",
			out: "usage: '$ crawl url-to-crawl' ",
		},
		{
			in:  "http://www.emergeadapt.com/",
			out: "begining crawl at: http://www.emergeadapt.com/",
		},
	}
	for i, test := range commandTests {
		actual := displayMsg(test.in)
		assert.Equal(t, test.out, actual, "Test %d", i)
	}
}

// Test crawl target URL
func TestRetrieveHTML(t *testing.T) {

}

// Test crawl rate limiting
func TestCrawlRequestDelay(t *testing.T) {

}

// Test link extraction from HTML
func TestExtractURLs(t *testing.T) {

}

// Test creating the output file
func TestFileWrite(t *testing.T) {

}
