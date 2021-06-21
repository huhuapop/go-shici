package handle

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
)

type AuthorHandle struct {}

func (h *AuthorHandle) Worker(body io.Reader, url string)   {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil{
		fmt.Errorf("doc.err", err)
	}
	doc.Find(".sons").Find(".cont").Find("a").Each(func(i int, selection *goquery.Selection) {
		authors :=selection.Text()
		fmt.Printf("%d author=%s\n", i, authors)
		link, _ := selection.Attr("href")
		fmt.Printf("%d link=%s\n", i, link)
	})
}
