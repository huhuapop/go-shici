package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/huhuapop/go_shici/handle"
	"log"
	"net/http"
)

func main()  {
	authors := "https://so.gushiwen.cn/authors/"
	//res, err := http.Get(authors)
	//if err != nil{
	//	panic(err)
	//}
	//
	//defer res.Body.Close()
	//if res.StatusCode != 200 {
	//	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	//}
	//
	//doc, err := goquery.NewDocumentFromReader(res.Body)
	//if err != nil{
	//	fmt.Errorf("doc.err", err)
	//}
	//doc.Find(".sons").Find(".cont").Find("a").Each(func(i int, selection *goquery.Selection) {
	//	authors :=selection.Text()
	//	fmt.Printf("%d author=%s\n", i, authors)
	//	link, _ := selection.Attr("href")
	//	fmt.Printf("%d link=%s\n", i, link)
	//})


	//fmt.Print("hello")
	h := handle.AuthorHandle{}
}