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
			out: "usage: $ go run ./main.go url-to-crawl",
		},
		{
			in:  "-h",
			out: "Available commands" + "'-h'  :   List available commands",
		},
		{
			in:  "12",
			out: "usage: $ go run ./main.go url-to-crawl",
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

// Test creating the output file
func TestWriteCSV(t *testing.T) {
	var filenameTests = []struct {
		path string
		file string
		out  string
	}{
		{
			path: ".",
			file: "output.csv",
			out:  "./output.csv",
		},
		{
			path: "..",
			file: "output.csv",
			out:  "../output.csv",
		},
		{
			path: ".",
			file: "out.put.csv",
			out:  "./out.put.csv",
		},
	}
	for i, test := range filenameTests {
		actual := writeCSV(test.path, test.file)
		assert.Equal(t, test.out, actual, "Test %d", i)
	}

}

func TestStoreResponse(t *testing.T) {

}

// Test link extraction from HTML
func TestScrapeLinks(t *testing.T) {

}

func TestRemoveIndex(t *testing.T) {

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

func TestIsTitle(t *testing.T) {

}

// Test crawl target URL
func TestGetHtmlTitle(t *testing.T) {

}

func TestWalkHTML(t *testing.T) {

}
