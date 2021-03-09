package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://avalith.net",
	"https://github.com/jonathanbs9/",
	"https://github.com/",
	"https://clockify.me/tracker",
	"https://www.youtube.com",
	"https://www.facebook.com",
	"https://www.twitter.com",
	"https://www.instagram.com/jonathanbs86",
}

func checkStatus(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for _ , url := range urls {
		wg.Add(1)
		go func (url string){
			resp, err := http.Get(url)
			if err != nil{
				fmt.Fprintf(w, "%+v \n ", err)
			}
			fmt.Fprintf(w, " url => %s %+v \n ", url, resp.Status)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	fmt.Println("Go waitGroup Example (main)")

	http.HandleFunc("/", checkStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Finished executing main program (main)")
}
