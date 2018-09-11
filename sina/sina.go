package main

import (

	"log"
	"github.com/PuerkitoBio/goquery" // 解析html
	//"io/ioutil"

)

func getAllUrls()  {
	var url string
	url = "https://weibo.com/3172815517/GyyY95V2R?filter=hot&root_comment_id=4282375879285384&type=comment#_rnd1536487962614"
	doc, err := goquery.NewDocument(url)      //获取将要爬取的html文档信息
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(" .WB_text > a").Each(func(i int, s *goquery.Selection) {
		html_url ,_:= s.Attr("href")
		log.Println(html_url)
	})
}



// 下载图片


func main()  {
	getAllUrls()

}