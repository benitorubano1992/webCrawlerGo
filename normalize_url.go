package main

import (
	"fmt"
	"net/url"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := parsedURL.Host + parsedURL.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}


 func getHeadingFromHTML(html string) string{
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil{
		return ""
	}
	headingSelection:= doc.Find("h1")
	if headingSelection.Length() == 0{
		headingSelection =doc.Find("h2")
	}

	if headingSelection.Length() == 0{
		return ""
	}
	return strings.TrimSpace(headingSelection.Text())

}

func getFirstParagraphFromHTML(html string) string{
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil{
		return ""
	}
	MainSelection:= doc.Find("main")
	if MainSelection.Length() == 1{
		pSelection:= MainSelection.Find("p")
		if pSelection.Length() > 0{
			return pSelection.First().Text()
		}
	}
	pSelect:= doc.Find("p")
	if pSelect.Length() == 0{
		return ""
	}
	return strings.TrimSpace(pSelect.First().Text())

}