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
	"errors"
	"net/url"
)

var redirectCount int = 0

func myRedirect(req *http.Request, via []*http.Request) (e error) {
	redirectCount++
	if redirectCount == 10 {
		redirectCount = 0
		return errors.New(req.URL.String())
	}
	return
}

func Extract(link string, keys string) ([]string, error) {
	var links []string
	client := &http.Client{
		CheckRedirect: myRedirect,
	}//解决重定向多次问题
	resp, err := client.Get(link)
	if err != nil {
		if e, ok := err.(*url.Error); ok && e.Err != nil {
			links = append(links, e.URL)
		}
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", link, resp.Status)
	}
	byte, _ := ioutil.ReadAll(resp.Body)
	s := string(byte)
	doc, err := html.Parse(bytes.NewReader(byte))
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", link, err)
	}
	resp.Body.Close()
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
			t := util.Substr(s, start, end - start)
			str := util.Str_delete(t)
			if len(str) > 0 {
				title = str
			}
		}
		db.Insert(keys, title, link)
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