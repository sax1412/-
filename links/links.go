package links

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"strings"
	_"github.com/go-sql-driver/mysql"
	"io/ioutil"
	"bytes"
	"../util"
	"../db"
)

func Extract(url string, keys string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	byte, _ := ioutil.ReadAll(resp.Body)
	s := string(byte)
	doc, err := html.Parse(bytes.NewReader(byte))
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	resp.Body.Close()
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				if strings.Contains(a.Val, "javascript") {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	if strings.Contains(s, keys) {
		title := "暂无标题"
		start, end := 0, 0
		h_index := []string{"<h1", "<h2", "<h3", "<h4", "<h5"}
		h_position := []string{"</h1>", "</h2>", "</h3>", "</h4>", "</h5>"}
		for k, v := range h_index {
			if strings.Contains(s, v) {
				start = strings.Index(s, v)
				end = strings.Index(s, h_position[k])
				break
			}
		}
		if end > 0 {
			title = util.Substr(s, start, end - start)
			title_slice := strings.Split(title, ">")
			title = title_slice[1]
		}
		db.Insert(keys, title, url)
	}
	forEachNode(doc, visitNode)
	return links, nil
}

func forEachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
}