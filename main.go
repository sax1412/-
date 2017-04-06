package main

import (
	"fmt"
	"log"
	"./links"
	"net/http"
	"os"
	"encoding/json"
)

func crawl(url string) []string {
	list, err := links.Extract(url, os.Args[2:3][0])
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
	worklist := make(chan []string, 20)
	var n = 1
	go func() {
		worklist <- os.Args[1:2]
	}()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
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