package api

import (
	"GiveMeMail/src/deal"
	"GiveMeMail/src/global"
	"fmt"
	"github.com/antchfx/htmlquery"
	"net/http"
	"net/url"
)

type SkymemApi struct {
	Name string
}

func NewSkymemApi() *SkymemApi{
	return &SkymemApi{"Skymem.info"}
}

//return name of this api
func (api SkymemApi) GetName() string{
	return api.Name
}

//main func of this api
func (api SkymemApi) Run() []string{
	return ReqAndParseSkymemApi()
}

//send a request to www.skymem.info ,parse and get the email info
//向www.skymem.info发送请求，解析并获得email信息
//params: domain:the email domain
//return value: the result slice of email
func ReqAndParseSkymemApi() []string{
	//init
	ParseResult :=make([]string,0)
	client:=deal.ProxyClient()

	//HTTP request
	DirectUrl:=GetDirectUrl()
	if DirectUrl==""{
		return ParseResult
	}
	UrlDericet,err:=url.Parse("https://www.skymem.info"+DirectUrl)
	if err!=nil{
		panic("error in parse Url")
	}

	req,_:=http.NewRequest("GET",UrlDericet.String(),nil)
	req.Header.Add("User-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36")

	//request repeatedly
	resp:=deal.RequestMore(req,client,5)
	if resp==nil{
		return ParseResult
	}


	//define a Xpath to find email
	XpathSample:="//tbody/tr[%d]/td[2]/a"

	//parse
	nodes,err:=htmlquery.Parse(resp.Body)
	if err!=nil{
		panic("error in parse html")
	}
	for i:=1;;i++{
		Xpath:=fmt.Sprintf(XpathSample,i)
		DivResult:=htmlquery.FindOne(nodes,Xpath)
		if DivResult==nil{
			break
		}
		ParseResult=append(ParseResult,htmlquery.InnerText(DivResult))

	}
	return ParseResult
}

//Because of the particularity of the website
//we need to get the link of the complete information page in this page first
//因为网站的特殊性，我们需要现在此网页中先获得完整信息页面的链接
//params: domian:the email domain
//return value: the url of complete information page
func GetDirectUrl() string{
	Domain:=global.Domain

	client:=deal.ProxyClient()

	//parse the url 解析url
	Url,err:=url.Parse("https://www.skymem.info/srch?ss=home&q="+Domain)
	if err!=nil{
		panic("parse url error!")
	}

	//make http request,and return the response  发起请求并返回结果
	req,_:=http.NewRequest("GET",Url.String(),nil)
	req.Header.Add("User-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36")
	resp := deal.RequestMore(req,client,5)
	if err!=nil{
		return ""
	}

	XpathDirect:="//div[@class='col-sm-7']/div/a"

	//parse the response
	//解析响应
	nodes, err := htmlquery.Parse(resp.Body)

	UrlDirect:=htmlquery.SelectAttr(htmlquery.FindOne(nodes,XpathDirect),"href")

	return UrlDirect

}