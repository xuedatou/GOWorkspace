package main

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery" // 解析html
	//"io/ioutil"
	"net/http"
	"os"
	"github.com/satori/go.uuid" // 生成图片文件名
	"io/ioutil"
	"golang.org/x/text/transform"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func getAllUrls() []string {

	var urls []string
	var url string
	//for i := 1246; i <10000; i++ {
	//	url = "https://dd.flexui.win/htm_data/8/1809/327"+strconv.Itoa(i)+".html"     //网址信息
	//
	//}
	url = ""
	doc, err := goquery.NewDocument(url)      //获取将要爬取的html文档信息
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".tr3 >.tal> h3 > a").Each(func(i int, s *goquery.Selection) {
		html_url, _ := s.Attr("href")
		log.Println(html_url)
		url = "https://dd.flexui.win/"+html_url;
		// 启动协程下载图片
		urls = append(urls, url)
	})

	return urls
}

func parseHtml(url string)  {
	doc, err := goquery.NewDocument(url)      //获取将要爬取的html文档信息
	if err != nil {
		log.Fatal(err)
	}
	p:=make(chan string)   //新开管道
	var fp string
	fp=doc.Find("tr >td > h4 ").Text()   //遍历整个文档



		log.Println(fp)
		fpps,_ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(fp)), simplifiedchinese.GBK.NewDecoder()))
		fp = string(fpps)
		log.Println(fp)
		os.Mkdir("G://xc//"+fp,os.ModePerm)


	doc.Find(".tr3 >td >p > b > b > b > input").Each(func(i int, s *goquery.Selection) {    //遍历整个文档

	//doc.Find("tr >td>.tpc_content > input").Each(func(i int, s *goquery.Selection) {
		img_url, _ := s.Attr("data-src")
		log.Println(img_url)


		// 启动协程下载图片
		go download(img_url,p,fp)        //将管道传入download函数
		fmt.Println("src = "+ <-p+"图片爬取完毕")
	})
	doc.Find("tr >td>.tpc_content > input").Each(func(i int, s *goquery.Selection) {
		img_url, _ := s.Attr("data-src")
		log.Println(img_url)


		// 启动协程下载图片
		go download(img_url,p,fp)        //将管道传入download函数
		fmt.Println("src = "+ <-p+"图片爬取完毕")
	})

}

// 下载图片
func download(img_url string,p chan string,fp string)  {
	uid, _ := uuid.NewV4()               //随机生成四段文件名
	file_name := uid.String() + ".jpg"
	fmt.Println(file_name)
	f, err := os.Create("G://xc//"+fp+"//"+file_name)
	if err != nil{
		log.Panic("文件创建失败")
	}
	defer f.Close()       //结束关闭文件

	resp, err := http.Get(img_url)
	if err != nil{
		fmt.Println("http.get err",err)
	}

	body,err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil{
		fmt.Println("读取数据失败")
	}
	defer resp.Body.Close()     //结束关闭
	f.Write(body)
	p <- file_name    //将文件名传入管道内

}

func main()  {
	 urls :=getAllUrls()
	for _, url := range urls {
		parseHtml(url )
	}
}