package main

import (
"fmt"
"os"
"strings"
"golang.org/x/net/html"
)

var raw =`<!DOCTYPE html>
<html>
<body>
<h1>My First Heading</h1>
<p>My first paragraph.</p>
<p>HTML images are defined with the img tag:</p>
<img src="xxx.jpg" width="104" height="142">
</body>
</html>`


func visit(n *html.Node, words *int, images *int) {
	// if it's an element node then see what tag it has
	fmt.Printf("type:%v , data:%v %v \n",n.Type , n.Data, n.Type == html.TextNode);

	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, images)
	}
}


func countWordsAndPics(n *html.Node) (int, int) {	
	var words, images int
	visit(n, &words, &images)
	return words, images
}


func main() {
	doc, err := html.Parse(strings.NewReader(raw))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing HTML: %v\n", err)
		return
	}

	words , pics := countWordsAndPics(doc)
	fmt.Printf("Words: %d, Pictures: %d\n", words, pics)
}