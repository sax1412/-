package main

import (
	"fmt"
	"log"
	"./links"
	"net/http"
	"encoding/json"
	"os"
)

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func kid(w http.ResponseWriter, r *http.Request) {
	kid := map[string]interface{}{}
	kid["kid"] = "jide"
	new_str, _ := json.Marshal(kid)
	w.Write([]byte(new_str))
	fmt.Println(w)
}

func main() {
	worklist := make(chan []string, 10)
	go func() {
		worklist <- os.Args[1:]
	}()
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
	//http.HandleFunc("/123", kid)
	//http.ListenAndServe(":4000", nil)
	//res, _ := http.Get("http://jinyun.datahunter.cn/api/ok")
	//body, _ := ioutil.ReadAll(res.Body)
	//jsons := map[string]interface{}{}
	//json.Unmarshal(body, &jsons)
	//fmt.Println(jsons)
}