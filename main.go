package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"golang.org/x/net/html"
	"bytes"
	"net/url"
)

func main(){
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://httpbin.org/status/404",
	}
	sucessCode :=0
	for i:=0;i<len(urls);i++{
		if checkURL(urls[i]){
			sucessCode++
		}
	}	
	fmt.Println("Sucessful Code :",sucessCode)
	fmt.Println("Finish Checking URL")
}

func checkURL(url string)bool{
	resp,err := http.Get(url)
		if err != nil {
			fmt.Println("Error :",err)
			return false
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Read Error :",err)
			return false
		}

		if strings.Contains(string(body),"Example Domain"){
			fmt.Println("Found Example Domain")
		}
		links := extractLinks(body)

		fmt.Println("Links found:", len(links))
		for _, link := range links {
    		fullURL, err := resolveURL(url, link)
   			if err != nil {
        		fmt.Println("URL Error:", err)
        		continue
    		}

    		fmt.Println(fullURL)
		}

			fmt.Println("HTML Length :",len(body))
		fmt.Println("URL : ",url)
		fmt.Println("Status Code : ",resp.StatusCode)	
		return resp.StatusCode == 200
}



func extractLinks(body []byte) []string {
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Println("HTML Parse Error:", err)
		return nil
	}

	var links []string

	var findLinks func(*html.Node)
	findLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			findLinks(child)
		}
	}

	findLinks(doc)
	return links
}

func resolveURL(base string, href string) (string, error) {
    baseURL, err := url.Parse(base)
    if err != nil {
        return "", err
    }

    hrefURL, err := url.Parse(href)
    if err != nil {
        return "", err
    }

    return baseURL.ResolveReference(hrefURL).String(), nil
}

