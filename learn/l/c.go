package main

import (
	"context"

	//"fmt"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 连接我远程服务器上启动和chrome-headless 服务器
	// 因为我的代码不是在我的笔记本上运行,所以不能使用client.New默认配置
	// 所以使用client.URL来自定义自己服务器地址
	c, err := chromedp.New(ctxt, chromedp.WithTargets(client.New(client.URL("http://pan.mojotv.cn:9222/json")).WatchPageTargets(ctxt)), chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var siteHref, title, iFrameCode string
	err = c.Run(ctxt, visitMojoTvDotCn("https://mojotv.cn/2018/12/10/how-to-create-a-https-proxy-serice-in-100-lines-of-code.html", &siteHref, &title, &iFrameCode))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("`%s` (%s),html:::%s", title, siteHref, iFrameCode)
}

func visitMojoTvDotCn(url string, elementHref, pageTitle, iFrameHtml *string) chromedp.Tasks {
	//临时放图片buf
	var buf []byte
	return chromedp.Tasks{
		//跳转到页面
		chromedp.Navigate(url),
		//chromedp.Sleep(2 * time.Second),
		//等待博客正文显示
		chromedp.WaitVisible(`#post`, chromedp.ByQuery),
		//滑动页面到google adsense 广告
		chromedp.ScrollIntoView(`ins`, chromedp.ByQuery),
		chromedp.Screenshot(`#post`, &buf, chromedp.ByQuery, chromedp.NodeVisible),
		//截图到文件
		chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
			return ioutil.WriteFile("mojotv.png", buf, 0644)
		}),
		//等待mojotv google广告展示出来
		chromedp.WaitVisible(`ins`, chromedp.ByQuery),
		//获取我的google adsense 广告代码
		chromedp.InnerHTML(`ins`, iFrameHtml, chromedp.ByQuery),
		//跳转到我的bilibili网站
		chromedp.Click("#copyright > a:nth-child(3)", chromedp.NodeVisible),
		//等待则个页面显现出来
		chromedp.WaitVisible(`#page-index`, chromedp.ByQuery),
		//在chrome浏览器页面里执行javascript
		chromedp.Evaluate(`document.title`, pageTitle),
		chromedp.Screenshot(`#page-index`, &buf, chromedp.ByQuery, chromedp.NodeVisible),
		//截取bili网页图片
		chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
			return ioutil.WriteFile("bili.png", buf, 0644)
		}),
		//获取bilibili网页的标题
		chromedp.JavascriptAttribute(`a`, "href", elementHref, chromedp.ByQuery),
	}
}