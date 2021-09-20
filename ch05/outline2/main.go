package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// forEachNode は n から始まるツリー内の個々のノード x に対して
// 関数 pre(x) と post(x) を呼び出します。その2つの関数はオプションです。
// pre は子ノードを訪れる前に呼び出され（前順: preorder）
// post は子ノードを訪れた後に呼び出されます（後順: postorder）
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Printf("getting %s: %s\n", url, resp.Status)
			return
		}

		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("parsing %s as HTML: %v\n", url, err)
			return
		}

		forEachNode(doc, startElement, endElement)
	}
}
