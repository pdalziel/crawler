package main

import (
	"testing"
	//"net/http/httptest"
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
			out: "usage: $ go run ./main.go url-to-craw",
		},
		{
			in:  "-h",
			out: "Available commands" + "'-h'  :   List available commands",
		},
		{
			in:  "12",
			out: "usage: $ go run ./main.go url-to-craw",
		},

	}
	for i, test := range commandTests {
		actual := displayMsg(test.in)
		assert.Equal(t, test.out, actual, "Test %d", i)
	}
}

func TestListCommands(t *testing.T) {
	actual := listCommands()
	assert.Equal(t, "Available commands"+"'-h'  :   List available commands", actual)

}


// Test crawl target URL
func TestGetHtmlTitle(t *testing.T) {

}

// Test link extraction from HTML
func TestScrapeLinks(t *testing.T) {

}

// Test creating the output file
func TestWriteCSV(t *testing.T) {

}

func TestRemoveIndex(t *testing.T) {

}

func TestIsTitle(t *testing.T) {

}

func TestWalkHTML(t *testing.T) {

}

func TestStoreResponse(t *testing.T) {

}

func TestStoreLinks(t *testing.T) {

}

func TestCleanLinks(t *testing.T) {

}

func TestCheckDomain(t *testing.T) {

}

func TestScrapeAll(t *testing.T) {

}

func TestEnqueue(t *testing.T) {

}
