package main 

import (
	"net/url"
	"fmt"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages int
}






/*func(cfg *config) sameDomain(currentUrl string)(bool,error){
	
	parsedCurrentUrl,err:= url.Parse(currentUrl)
	if err != nil {
		return false, fmt.Errorf("couldn't parse currentUrl:%s,err: %w", currentUrl,err)
	}
	return cfg.baseURL.HostName() == parsedCurrentUrl.HostName(),nil

}*/


func(cfg *config) canCrawlMore()bool{
	cfg.mu.Lock()
	//fmt.Printf("cfg numPagesFetch:%d, maxPagesFetch:%d\n",len(cfg.pages),cfg.maxPages)

	defer cfg.mu.Unlock()
	return len(cfg.pages) < cfg.maxPages

}


func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool){
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	_,exists:= cfg.pages[normalizedURL]
	if !exists{
		cfg.pages[normalizedURL]=PageData{
			URL:normalizedURL,
		}
	}
	return exists == false
}

func (cfg *config) addPageData(url string,p PageData){
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[url]=p
}



func(cfg *config)crawlPage(rawCurrentURL string){
	if !cfg.canCrawlMore(){
		cfg.wg.Done()
		return
	}
	
	cfg.concurrencyControl<-struct{}{}
	defer func(){
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	




	parsedCurrentUrl,err:= url.Parse(rawCurrentURL)
	if err != nil {
		 fmt.Printf("couldn't parse currentUrl:%s,err: %w", rawCurrentURL,err)
	}
	if cfg.baseURL.Hostname() != parsedCurrentUrl.Hostname(){
		fmt.Printf("baseURL: %s, currentUrl:%s has different HostName\n", cfg.baseURL.String(),rawCurrentURL)
		return
	}

	
	
	nUrlCurrent,err:=normalizeURL(rawCurrentURL)
	if err != nil{
		fmt.Printf("err Normalize Url:%s,err:%v\n",rawCurrentURL,err)
		return
	}
	hasBeenVisited:=cfg.addPageVisit(nUrlCurrent)
	if !hasBeenVisited{
		return
	}
	hmtlStr,err:= getHTML(rawCurrentURL)
	if err != nil{
		fmt.Printf("Err get Html url: %s, err:%v\n",rawCurrentURL,err)
		return
	}
	

	//fmt.Printf("Get htmlString from url :%s\n",rawCurrentURL)
	pageData,err:=extractPageData(hmtlStr,rawCurrentURL)
	
	if err != nil{
		fmt.Printf("error retrieving links from html:%s\nurl:%s, err:%v\n",hmtlStr,rawCurrentURL,err)
		return
	}
	cfg.addPageData(rawCurrentURL,pageData)
	fmt.Printf("url:%s, has spawn more %d goroutine\n",rawCurrentURL,len(pageData.OutgoingLinks))

	for _,link:= range pageData.OutgoingLinks{
		
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}

	/*urlLinks,err:= getURLsFromHTML(hmtlStr,parsedCurrentUrl)
	if err != nil{
		fmt.Printf("error retrieving links from html:%s\nurl:%s, err:%v\n",hmtlStr,nUrlCurrent,err)
		return
	}

	for link:= range urlLinks{
		cfg.concurrencyControl<-struct{}{}
		wg.Add(1)
		go cfg.crawlPage(link)
	}*/

	

}


/*func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int){
	hasSameDomain,err:= sameDomain(rawBaseURL,rawCurrentURL)
	if err != nil{
		fmt.Printf("Error :%v\n",err)
		return
	}
	if !hasSameDomain{
		fmt.Printf("baseURL: %s, currentUrl:%s has different HostName\n", rawBaseURL,rawCurrentURL)
		return
	}
	nUrlCurrent,err:=normalizeURL(rawCurrentURL)
	if err != nil{
		fmt.Printf("err Normalize Url:%s,err:%v\n",rawCurrentURL,err)
		return
	}

	page,exists:= pages[nUrlCurrent]
	if exists{
		pages[nUrlCurrent]=page + 1
		return
	}
	pages[nUrlCurrent]=1
	fmt.Printf("Calling getHtml url:%s\n",nUrlCurrent)
	hmtlStr,err:= getHTML(rawCurrentURL)
	if err != nil{
		fmt.Printf("Err get Html url: %s, err:%v\n",nUrlCurrent,err)
		return
	}
	pageData,err:=extractPageData(hmtlStr,rawCurrentURL)
	if err != nil{
		fmt.Printf("error retrieving links from html:%s\nurl:%s, err:%v\n",hmtlStr,nUrlCurrent,err)
		return
	}
	for _,link:= range pageData.OutgoingLinks{
		crawlPage(rawBaseURL,link,pages)
	}
	
}
*/