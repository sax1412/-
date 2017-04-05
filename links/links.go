package links

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"strings"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
	"io/ioutil"
	"bytes"
)

var db, _ = sql.Open("mysql", "root@/pachong?charset=utf8")

func insert(name string, url string) {
	exe, _ := db.Prepare("insert star set name = ?,url = ?, ct = ?")
	exe.Exec(name, url, time.Now())
}

func Extract(url string) ([]string, error) {
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
	doc, _ := html.Parse(bytes.NewReader(byte))
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
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
	if strings.Contains(s, "范冰冰") {
		insert("范冰冰", url)
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