package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

)

type htmlStruct struct {
	Body body `xml:"body"`
}
type body struct {
	Content string `xml:",innerxml"`
}

const testAnchorsHTML = `
<html>
   <body>
   		<div>
   			<p id="a" class="t1">
   				not an anchor!
   				<br/>
   				<a class="t3">test anchor</a>
   			</p>
   			<p id="b" class="t2">bbb
   				<a class="t1">another anchor</a>
   			</p>
   			<p id="c" class="t3">ccc</p>
   			<p id="d" class="t4">ddd</p>
   		</div>
   	</body>
 </html>
 `

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
	}
	for i, test := range filenameTests {
		actual := writeCSV(test.path, test.file)
		assert.Equal(t, test.out, actual, "Test %d", i)
	}

}

// TODO refactor storeResponse to add map params
func TestStoreResponse(t *testing.T) {
	response := httpResponse{"https://www.emergeadapt.com", "Business Process & Case Management Platform in the Cloud- CaseBlocks", "200 OK"}
	m := map[string]httpResponse{"https://www.emergeadapt.com": response}
	fmt.Println(m)
	var storeTests = []struct {
		uri    string
		title  string
		status string
		out    string
	}{
		{
			uri:    "https://www.emergeadapt.com",
			title:  "Business Process & Case Management Platform in the Cloud- CaseBlocks",
			status: "200 OK",
			out:    "Not adding duplicate: https://www.emergeadapt.com",
		},
	}
	for i, test := range storeTests {
		actual := storeResponse(test.uri, test.title, test.status)
		assert.Equal(t, test.out, actual, "Test %d", i)
	}

}

//TODO
func TestScrapeLinks(t *testing.T) {

}

func TestRemoveIndex(t *testing.T) {
	expected := []string{"http://www.emergeadapt.com", "http://www.emergeadapt.com/about"}
	arr := []string{"http://www.emergeadapt.com", "http://www.emergeadapt.com/login",
		"http://www.emergeadapt.com/about"}
	actual := removeIndex(arr, 1)
	assert.Equal(t, expected, actual, "Test %d")
}

//TODO
func TestCleanLinks(t *testing.T) {

}

//TODO
func TestCheckDomain(t *testing.T) {

}

//TODO
func TestIsTitle(t *testing.T) {
	b := []byte(`<!DOCTYPE html>
<html>
    <head>
        <title>
            Title of the document
        </title>
    </head>
    <body>
        body content
        <p>more content</p>
    </body>
</html>`)
	fmt.Println(b)

}

//TODO
func TestGetHTMLTitle(t *testing.T) {

}

//TODO
func TestWalkHTML(t *testing.T) {

}
