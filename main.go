package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/bobesa/go-domain-util/domainutil"

	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"
)

type link struct {
	uri     string
	scraped bool
}

type httpResponse struct {
	uri, title, response string
}

// collection of http responses
var m map[string]httpResponse

// collection of links
var linkMap map[string]link

// Display usage, flags and custom errors to stdout
func displayMsg(in string) string {
	switch in {
	case "":
		fmt.Println("usage: '$ crawl url-to-crawl'")
		return "usage: '$ crawl url-to-crawl'"
	case "-h":
		listCommands()
		// can be expanded
	default:
		fmt.Println("usage: '$ crawl url-to-crawl'")
	}
	return "usage: '$ crawl url-to-crawl'"
}

func listCommands() string {
	// Can be extended and used to list all available command flags
	fmt.Println("Available commands")
	fmt.Println("'-h'  :   List available commands")
	return "Available commands" + "'-h'  :   List available commands"
}

// Error logging
func logError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
		os.Exit(1) //DON'T PANIC
	}
}

// remove urls
func removeIndex(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// retrieve links scraped from html source of the target url
func scrapeLinks(uri string) []string {
	links := []string{}
	if link, ok := m[uri]; ok {
		fmt.Println("Not scraping duplicate: ", link)
	} else {
		resp, err := http.Get(uri)
		if err != nil {
			//logError("Cannot retrieve HTML: ", err)
			fmt.Println(uri, err)
			return links
		}
		title, _ := GetHtmlTitle(uri)
		// add http response to map
		storeResponse(uri, title, resp.Status)
		b := resp.Body
		defer resp.Body.Close()
		z := html.NewTokenizer(b)
		for {
			tt := z.Next()
			switch {
			case tt == html.ErrorToken:
				return links
			case tt == html.StartTagToken:
				t := z.Token()
				isAnchor := t.Data == "a"
				if isAnchor {
					links = append(links, t.String())
				}
			}
		}

	}

	return links
}

func isTitle(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func GetHtmlTitle(uri string) (string, bool) {
	resp, err := http.Get(uri)
	if err != nil {
		logError("Cannot retrieve HTML for title: ", err)
	}
	r := resp.Body
	doc, err := html.Parse(r)
	if err != nil {
		logError("Cannot parse title: ", err)
	}
	return walkHTML(doc)
}

func walkHTML(n *html.Node) (string, bool) {
	if isTitle(n) {
		return n.FirstChild.Data, true
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := walkHTML(c)
		if ok {
			return result, ok
		}
	}
	return "", false
}
func storeResponse(uri string, title string, status string) {
	if link, ok := m[uri]; ok {
		fmt.Println("Not adding duplicate: ", link)
	} else {
		m[uri] = httpResponse{
			uri,
			title,
			status,
		}
		fmt.Println("added: ", m[uri])
	}
}

func writeCSV() {
	var headers = []string{"url", "title", "status code"}
	data := []string{}
	for _, i := range m {
		data[0] = i.uri
		data[1] = i.title
		data[2] = i.response
		fmt.Println(i, i.title, i.response)
	}
	// setup writer
	csvOut, err := os.Create("http_responses_file.csv")
	if err != nil {
		log.Fatal("Unable to open output")
	}
	w := csv.NewWriter(csvOut)
	defer csvOut.Close()

	if err = w.Write(headers); err != nil {
		logError("Cannot write headers", err)
	}
	if err = w.Write(data); err != nil {
		logError("Cannot write data", err)
	}
	w.Flush()
}

// store the links in a set
func storeLinks(link string) {
	if uri, ok := linkMap[link]; ok {
		fmt.Println(uri, " is a duplicate of: ", link)
	} else {
		//linkMap[link]
	}
}

// remove html tags etc
func cleanLinks(links []string) []string {
	re := regexp.MustCompile("<a href=\"")
	reTail := regexp.MustCompile("\"")
	cleaned := []string{}
	for i := range links {
		value := links[i]
		result := re.ReplaceAllString(value, "")
		u := reTail.Split(result, -1)

		uStr := u[0]
		cleaned = append(cleaned, uStr)
	}
	return cleaned
}
func checkDomain(links []string, domain string) []string {
	for i := 0; i < len(links); i++ {
		if domainutil.Domain(links[i]) != domainutil.Domain(domain) {
			links = removeIndex(links, i)
			i-- // slice  is now one element shorter
		}
	}
	return links
}

func main() {
	m = make(map[string]httpResponse)
	linkMap = make(map[string]link)
	fmt.Println()
	var domain = "http://www.emergeadapt.com" // hard coded due to requirements
	flag.Parse()
	args := flag.Args()
	u, err := url.ParseRequestURI(args[0]) // check seed link is valid
	uri := u.String()
	fmt.Println("begining crawl at: " + uri)
	if err != nil {
		fmt.Println(u, err)
		displayMsg(args[0])
		logError("input error: "+args[0], err)
	}
	enqueue(uri, domain)
	// deferred calls are executed in last-in-first-out order
	defer writeCSV()
	for i := 0; i < len(linkMap); i++ {
		if uri, ok := m[linkMap[i]]; ok {
			fmt.Println(linkMap[i], " should never see this ... found duplicate: ", uri)
		} else {
			enqueue(linkMap[i], domain)
		}
	}
}

func scrapeAll(domain string) {
	fmt.Println(len(linkMap), " queued links")
	for i := 0; i < len(linkMap); i++ {
		if uri, ok := m[linkMap[i]]; ok {
			fmt.Println(linkMap[i], " should never see this ... found duplicate: ", uri)
		} else {
			enqueue(linkMap[i], domain)
		}
	}
}

func enqueue(uri string, domain string) {
	timer := time.NewTimer(time.Second * 3)
	<-timer.C // wait for 3 seconds
	// check if uri has already been visited
	if _, ok := m[uri]; ok {
		fmt.Println("should never see this ", uri)
	} else {
		links := scrapeLinks(uri)
		fmt.Println("scraping: ", uri)
		links = checkDomain(links, domain)
		cleanedLinks := cleanLinks(links)
		for i := 0; i < len(cleanedLinks); i++ {
			storeLinks(cleanedLinks[i])
		}
	}
}
