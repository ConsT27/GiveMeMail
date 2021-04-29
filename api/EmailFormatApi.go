package api

import (
	"GiveMeMail/src/deal"
	"GiveMeMail/src/global"
	"fmt"
	"github.com/antchfx/htmlquery"
	"net/http"
	"net/url"
	"strings"
)

type EmailFormatApi struct {
	Name string
}

func NewEmailFormatApi() *EmailFormatApi{
	return &EmailFormatApi{"Email-Format.com"}
}

//return name of this api
func (api EmailFormatApi) GetName() string{
	return api.Name
}

//main func of this api
func (api EmailFormatApi) Run() []string{
	ParseResult:=make([]string,0)
	ParseResult = reqAndParseEmailFormatApi()
	return ParseResult
}

//Send a request to www.email-format.com and parse the results to find email,then return them
//发送请求到www.email-format.com并解析其中的邮件内容返回
//params: domain:the email domain
func reqAndParseEmailFormatApi() []string{
	//init
	Domain:=global.Domain
	ParseResult :=make([]string,0)

	//HTTP Request
	client := deal.ProxyClient()

	//parse the url 解析url
	Url,err:=url.Parse("https://www.email-format.com/d/"+Domain)
	if err!=nil{
		panic("parse url error!")
	}

	//make http request,and return the response  发起请求并返回结果
	req,_:=http.NewRequest("GET",Url.String(),nil)
	req.Header.Add("User-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36")

	//request repeatedly
	resp:=deal.RequestMore(req,client,5)
	if resp==nil{
		return ParseResult
	}

	//define a Xpath to find email
	XpathSample:="//tr[%d]/td[@class='td_email']/div[@class='fl']"

	//parse the response
	//解析响应
	nodes, err := htmlquery.Parse(resp.Body)

	//parse the html with xpath and get the email info
	//用xpath解析html并获得email信息
	for i:=1;;i++ {
		Xpath:=fmt.Sprintf(XpathSample,i)
		DivInfo:=htmlquery.FindOne(nodes,Xpath)
		if DivInfo==nil{
			break
		}
		ParseResult=append(ParseResult,strings.TrimSpace(htmlquery.InnerText(DivInfo)))
	}
	return ParseResult
}

