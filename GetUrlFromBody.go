package main

import (
	"fmt"
	"net/url"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func getAbsoluteUrl(htmlUrl string,baseURL *url.URL)(string,error){
	parsedURL, err := url.Parse(htmlUrl)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	if parsedURL.Scheme != ""{
		return htmlUrl,nil
	}
	return baseURL.JoinPath(htmlUrl).String(),nil
}


func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	var result []string
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil{
		return result,err
	}
	bodySelection:= doc.Find("body")
	if bodySelection.Length() == 0{
		return result,nil
	}
	var errorResult error
	bodySelection.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
        // For each '<a href>' it finds, it will run this function.
		val, exists := s.Attr("href")
		if !exists{
			return
		}

		hrefText:= strings.TrimSpace(val)
		absPath,err:= getAbsoluteUrl(hrefText,baseURL)
		if err == nil{
			result=append(result,absPath)	
		}else if errorResult == nil && err != nil{
			errorResult=err
		}
		
		
	})
	if errorResult == nil{
		return result,nil
	}
	return []string{},errorResult

}


func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	var result []string
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil{
		return result,err
	}
	bodySelection:= doc.Find("body")
	if bodySelection.Length() == 0{
		return result,nil
	}
	
var errorResult error
	bodySelection.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
        // For each '<a href>' it finds, it will run this function.
		val, exists := s.Attr("src")
		if !exists{
			return
		}

		hrefText:= strings.TrimSpace(val)
		absPath,err:= getAbsoluteUrl(hrefText,baseURL)
		if err == nil{
			result=append(result,absPath)	
		}else if errorResult == nil && err != nil{
			errorResult=err
		}
		
		
	})
	if errorResult == nil{
		return result,nil
	}
	return []string{},errorResult
	
	


}