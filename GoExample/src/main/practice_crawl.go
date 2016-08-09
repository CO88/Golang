package main

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

var check_map map[string]bool = make(map[string]bool)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan string) {

    ch <- url
    if checkMap(url, ch) {
    	//fmt.Println("debug")
        return
    }
    
    if depth <= 0 {
        return
    }
    
    body, urls, err := fetcher.Fetch(url)
	

    if err != nil {
        fmt.Println(err)
        return
    } else {
		fmt.Printf("found: %s %q\n", url, body)
	}
    
    // done 채널을 이용하지 않으면 고루틴이 실행과 동시에 Crawl함수가 종료되면서
    // 의도치 않은 결과를 만들어 냄
    done := make(chan bool)
    for _, u := range urls {
		go func(url string) {
			Crawl(url, depth-1, fetcher, ch)
			done <- true
		}(u)
		
		<-done
    }
	
    return
}

func checkMap(url string, ch chan string) (ret bool){
    
	_ ,ok := check_map[url]
    
	if ok == false {
		check_map[url] = true
		ret = false
	} else {
		ret = true
	}	
	
    v := <- ch
	if v != url {
		panic(fmt.Sprintf("Should not be different?! : %v : %v !!!", v, url))
	}
    return
}


// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := (*f)[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
