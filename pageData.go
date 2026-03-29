package main
import (
	"net/url"
)

type PageData struct {
    URL            string   `json:"url"`
    Heading        string   `json:"heading"`
    FirstParagraph string   `json:"first_paragraph"`
    OutgoingLinks  []string `json:"outgoing_links"`
    ImageURLs      []string `json:"image_urls"`
}


func extractPageData(html, pageURL string) (PageData,error) {
	p:=PageData{}
	urlBase,err:= url.Parse(pageURL)
	if err != nil{
		return p,err
	}
	p.URL = urlBase.String()
	//urlS,err:= normalizeURL(pageURL)
	/*if err != nil{
		return p,err
	}
	p.URL = urlS
	*/
	p.Heading = getHeadingFromHTML(html)
	p.FirstParagraph = getFirstParagraphFromHTML(html)
	aLinks,err:= getURLsFromHTML(html,urlBase)
	if err != nil{
		return PageData{},err
	}
	imgLinks,err:= getImagesFromHTML(html,urlBase)
	if err != nil{
		return PageData{},err
	}
	p.OutgoingLinks = aLinks
	p.ImageURLs = imgLinks

	return p,nil



}