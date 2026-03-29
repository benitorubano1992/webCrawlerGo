package main
 import (
	"fmt"
	"os"
	"sync"
	"net/url"
	"strconv"
)




 

 


func main(){
	var wg sync.WaitGroup
	//fmt.Println("Hello, World!")
	args:=os.Args
	argsWitouthP:=args[1:]
	if len(argsWitouthP) == 0{
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(argsWitouthP) > 3{
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}


	urlBase:= argsWitouthP[0]
	maxConcurrency:= 3
	maxPages:= 20
	if len(argsWitouthP) >= 2{
		maXConcurr,err:= strconv.Atoi(argsWitouthP[1])
		if err != nil{
			fmt.Printf("maxConcurrency Value: %s is not a valid number\n",argsWitouthP[1])
		}else{
			maxConcurrency = maXConcurr
		}
	}
	if len(argsWitouthP) == 3{
		maxPagesV,err:= strconv.Atoi(argsWitouthP[2])
		if err != nil{
			fmt.Printf("maxPages Value: %s is not a valid number\n",argsWitouthP[2])
			os.Exit(1)
		}
		maxPages = maxPagesV
	}



	
	rawUrl,err:= url.Parse(urlBase)
	if err != nil {
		 fmt.Printf("couldn't parse currentUrl:%s,err: %w", urlBase,err)
		os.Exit(1)
	}
	//fmt.Printf("url Base: %s, maxConcurrency:%d, maxPages:%d\n",urlBase,maxConcurrency,maxPages)
	//fmt.Printf("starting crawl of: %s\n",urlBase)
	cfg:=config{
		pages:make(map[string]PageData),
		baseURL:rawUrl,
		concurrencyControl:make(chan struct{},maxConcurrency),
		wg:&wg,
		mu:&sync.Mutex{},
		maxPages:maxPages,
	}
	
	cfg.wg.Add(1)
	go cfg.crawlPage(urlBase)

	
	cfg.wg.Wait()
	/*for key,value:=range cfg.pages{
		fmt.Printf("url:%s has been visited, links:%+v\n",key,value)
	}*/
	if err:=writeJSONReport(cfg.pages, "report.json");err != nil{
		fmt.Printf("Err Writing :%v\n",err)
	}

	/*pages:=make(map[string]int)
	crawlPage(urlBase,urlBase,pages)
	*/
	/*if err != nil{
		fmt.Printf("Err: %v\n",err)
		os.Exit(1)
	}
	fmt.Println(resHtml)
	*/

}

