package main 
import (
	"net/http"
	"io"
	"fmt"
	"strings"
)

func getHTML(rawURL string) (string, error){
	req,err:= http.NewRequest("GET",rawURL,nil)
	if err != nil{
		return "",fmt.Errorf("getHtml Create Request url:%s, err:%w",rawURL,err)
	}
	req.Header.Set("User-Agent","BootCrawler/1.0")
	req.Header.Set("Content-Type","application/json")
	client:=&http.Client{}
	res,err:= client.Do(req)
	if err != nil{
		return "",fmt.Errorf("getHtml Get Response url:%s, err:%w",rawURL,err)
	}
	if res.StatusCode >= http.StatusBadRequest{
		return "",fmt.Errorf("getHtml Get url:%s, statusCode:%d, status:%s",rawURL,res.StatusCode,res.Status)
	}
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s", contentType)
	}
	defer res.Body.Close()
	resBodyBytes,err:= io.ReadAll(res.Body)
	if err != nil{
		return "",fmt.Errorf("getHtml Get url:%s, retrieving body: %w",rawURL,err)
	}

	return string(resBodyBytes),nil

}